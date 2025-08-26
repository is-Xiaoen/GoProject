package impl

import (
	"context"
	"time"

	"github.com/infraboard/mcube/v2/desense"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token"
)

// IssueToken 登录接口(颁发Token)
func (i *TokenServiceImpl) IssueToken(ctx context.Context, in *token.IssueTokenRequest) (*token.Token, error) {
	// 颁发Token
	// 根据请求中的颁发器类型（如 user/password, ldap, 飞书等），获取对应的颁发器实例
	issuer := token.GetIssuer(in.Issuer)
	// 如果颁发器不存在，则返回不支持的错误
	if issuer == nil {
		return nil, exception.NewBadRequest("issuer %s no support", in.Issuer)
	}

	// 调用对应颁发器的 IssueToken 方法，生成令牌
	tk, err := issuer.IssueToken(ctx, in.Parameter)
	if err != nil {
		return nil, err
	}

	// 设置令牌的颁发者和来源
	tk.SetIssuer(in.Issuer).SetSource(in.Source)

	// 判断当前数据库有没有已经存在的同源活跃令牌
	activeTokenQueryReq := token.NewQueryTokenRequest().
		AddUserId(tk.UserId).
		SetSource(in.Source).
		SetActive(true)

	tks, err := i.QueryToken(ctx, activeTokenQueryReq)
	if err != nil {
		return nil, err
	}

	// 根据令牌来源类型，决定如何处理已存在的活跃令牌
	switch in.Source {
	// 对于 Web、iOS、Android 和 PC 端，每个端只能有一个活跃登录会话
	case token.SOURCE_WEB, token.SOURCE_IOS, token.SOURCE_ANDROID, token.SOURCE_PC:
		if tks.Len() > 0 {
			// 如果已存在活跃令牌，则直接返回已有的令牌，而不是颁发新的
			i.log.Debug().Msgf("use exist active token: %s", desense.Default().DeSense(tk.AccessToken, "4", "3"))
			return tks.Items[0], nil
		}
	// 对于 API 令牌，限制活跃令牌数量
	case token.SOURCE_API:
		if tks.Len() > int(i.MaxActiveApiToken) {
			// 如果活跃 API 令牌数量超过最大限制，则返回错误
			return nil, exception.NewBadRequest("max active api token overflow")
		}
	}

	// 如果没有已存在的活跃令牌（或策略允许），则将新生成的令牌保存到数据库
	if err := datasource.DBFromCtx(ctx).
		Create(tk).
		Error; err != nil {
		return nil, err
	}
	return tk, nil
}

// ValiateToken 校验Token 是给内部中间层使用 身份校验层
func (i *TokenServiceImpl) ValiateToken(ctx context.Context, req *token.ValiateTokenRequest) (*token.Token, error) {
	// 1. 查询Token (检查令牌是否由本系统颁发)
	tk := token.NewToken()
	err := datasource.DBFromCtx(ctx).
		Where("access_token = ?", req.AccessToken).
		First(tk).
		Error
	if err != nil {
		// 如果查询不到，返回错误
		return nil, err
	}

	// 2.1 判断访问令牌（AccessToken）是否过期
	if err := tk.IsAccessTokenExpired(); err != nil {
		// 如果访问令牌过期，继续判断刷新令牌（RefreshToken）是否过期
		if err := tk.IsRreshTokenExpired(); err != nil {
			// 如果刷新令牌也过期，则无法刷新，返回错误
			return nil, err
		}
		// 如果开启了自动刷新功能，则更新令牌的过期时间
		if i.AutoRefresh {
			tk.SetRefreshAt(time.Now())
			tk.SetExpiredAtByDuration(i.refreshDuration, 4)
			// 保存更新后的令牌状态
			if err := datasource.DBFromCtx(ctx).Save(tk); err != nil {
				i.log.Error().Msgf("auto refresh token error, %s", err.Error)
			}
		}
		// 即使刷新成功，由于访问令牌已过期，仍然返回原始的过期错误，让调用方知道需要用刷新令牌重新获取
		return nil, err
	}

	// 如果访问令牌有效且未过期，则返回令牌信息
	return tk, nil
}

// DescribeToken 查询Token详情
func (i *TokenServiceImpl) DescribeToken(ctx context.Context, in *token.DescribeTokenRequest) (*token.Token, error) {
	// 获取数据库上下文
	query := datasource.DBFromCtx(ctx)
	// 根据请求的描述方式构建查询条件
	switch in.DescribeBy {
	case token.DESCRIBE_BY_ACCESS_TOKEN:
		query = query.Where("access_token = ?", in.DescribeValue)
	default:
		// 如果描述方式不支持，则返回错误
		return nil, exception.NewBadRequest("unspport describe type %s", in.DescribeValue)
	}

	// 创建一个空的令牌对象用于接收查询结果
	tk := token.NewToken()
	// 执行查询
	if err := query.First(tk).Error; err != nil {
		return nil, err
	}
	return tk, nil
}

// RevolkToken 退出接口(销毁Token)
func (i *TokenServiceImpl) RevolkToken(ctx context.Context, in *token.RevolkTokenRequest) (*token.Token, error) {
	// 1. 描述令牌，通过访问令牌获取其详情
	tk, err := i.DescribeToken(ctx, token.NewDescribeTokenRequest(in.AccessToken))
	if err != nil {
		return nil, err
	}

	// 2. 校验传入的刷新令牌是否与数据库中的匹配
	if err := tk.CheckRefreshToken(in.RefreshToken); err != nil {
		return nil, err
	}

	// 3. 锁定令牌，设置锁定类型和原因
	tk.Lock(token.LOCK_TYPE_REVOLK, "user revolk token")

	// 4. 更新数据库中令牌的状态，使其失效
	err = datasource.DBFromCtx(ctx).Model(&token.Token{}).
		Where("access_token = ?", in.AccessToken).
		Where("refresh_token = ?", in.RefreshToken).
		Updates(tk.Status.ToMap()).
		Error
	if err != nil {
		return nil, err
	}
	return tk, err
}

// QueryToken 查询已经颁发出去的Token
func (i *TokenServiceImpl) QueryToken(ctx context.Context, in *token.QueryTokenRequest) (*types.Set[*token.Token], error) {
	// 创建一个结果集对象
	set := types.New[*token.Token]()
	// 获取数据库上下文
	query := datasource.DBFromCtx(ctx).Model(&token.Token{})

	// 根据请求参数构建动态查询条件
	if in.Active != nil {
		if *in.Active {
			// 如果查询活跃令牌，则条件为：未被锁定且刷新令牌未过期
			query = query.
				Where("lock_at IS NULL AND refresh_token_expired_at > ?", time.Now())
		} else {
			// 如果查询不活跃令牌，则条件为：已被锁定或刷新令牌已过期
			query = query.
				Where("lock_at IS NOT NULL OR refresh_token_expired_at <= ?", time.Now())
		}
	}
	if in.Source != nil {
		query = query.Where("source = ?", *in.Source)
	}
	if len(in.UserIds) > 0 {
		query = query.Where("user_id IN ?", in.UserIds)
	}

	// 1. 查询满足条件的总量
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	// 2. 执行分页查询，获取具体数据
	err = query.
		Order("issue_at desc").          // 按颁发时间倒序排序
		Offset(int(in.ComputeOffset())). // 计算并设置分页偏移量
		Limit(int(in.PageSize)).         // 设置分页大小
		Find(&set.Items).                // 查找结果并填充到 Items 中
		Error
	if err != nil {
		return nil, err
	}
	return set, nil
}

// 用户切换空间
// 该函数目前被注释掉，但其逻辑是根据用户ID和命名空间ID查询权限，如果用户有权限，则更新其令牌中的命名空间信息并保存。
// func (i *TokenServiceImpl) ChangeNamespce(ctx context.Context, in *token.ChangeNamespceRequest) (*token.Token, error) {
// 	set, err := i.policy.QueryNamespace(ctx, policy.NewQueryNamespaceRequest().SetUserId(in.UserId).SetNamespaceId(in.NamespaceId))
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	ns := set.First()
// 	if ns == nil {
// 		return nil, exception.NewPermissionDeny("你没有该空间访问权限")
// 	}
//
// 	// 更新Token
// 	tk, err := i.DescribeToken(ctx, token.NewDescribeTokenRequest(in.AccessToken))
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	tk.NamespaceId = ns.Id
// 	tk.NamespaceName = ns.Name
// 	// 保存状态
// 	if err := datasource.DBFromCtx(ctx).
// 		Updates(tk).
// 		Error; err != nil {
// 		return nil, err
// 	}
//
// 	return tk, nil
// }
