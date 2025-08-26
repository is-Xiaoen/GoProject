package privatetoken

import (
	"context"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/user"
)

// init 函数是 Go 语言的特殊函数，在包被导入时自动执行
func init() {
	// 将 PrivateTokenIssuer 实例注册到 ioc 配置中心
	ioc.Config().Registry(&PrivateTokenIssuer{})
}

// PrivateTokenIssuer 实现了 token.Issuer 接口，用于通过已有的访问令牌颁发新的令牌
type PrivateTokenIssuer struct {
	ioc.ObjectImpl // 嵌入 ioc.ObjectImpl，提供基础的生命周期管理能力

	user  user.Service  // 用户服务接口，用于获取用户信息
	token token.Service // 令牌服务接口，用于校验令牌
}

// Name 返回颁发器的名称，用于在 ioc 容器中注册
func (p *PrivateTokenIssuer) Name() string {
	return "private_token_issuer"
}

// Init 是 ioc 容器在服务启动时调用的初始化方法
func (p *PrivateTokenIssuer) Init() error {
	// 从 ioc 容器中获取用户服务和令牌服务实例
	p.user = user.GetService()
	p.token = token.GetService()
	// 将自身注册为 token 包中的颁发器，并指定名称
	token.RegistryIssuer(token.ISSUER_PRIVATE_TOKEN, p)
	return nil
}

// IssueToken 实现了 token.Issuer 接口的 IssueToken 方法
func (p *PrivateTokenIssuer) IssueToken(ctx context.Context, parameter token.IssueParameter) (*token.Token, error) {
	// 1. 校验传入的旧令牌是否合法
	oldTk, err := p.token.ValiateToken(ctx, token.NewValiateTokenRequest(parameter.AccessToken()))
	if err != nil {
		return nil, err
	}

	// 2. 查询令牌对应的用户
	uReq := user.NewDescribeUserRequestById(oldTk.UserIdString())
	u, err := p.user.DescribeUser(ctx, uReq)
	if err != nil {
		// 如果用户不存在，则返回“未授权”错误
		if exception.IsNotFoundError(err) {
			return nil, exception.NewUnauthorized("%s", err)
		}
		return nil, err
	}

	// 检查该用户是否开启了 API 登录权限
	if !u.EnabledApi {
		return nil, exception.NewPermissionDeny("未开启接口登录")
	}

	// 3. 颁发新令牌
	tk := token.NewToken()
	// 设置新令牌关联的用户信息
	tk.UserId = u.Id
	tk.UserName = u.UserName
	tk.IsAdmin = u.IsAdmin

	// 获取请求中指定的过期时间
	expiredTTL := parameter.ExpireTTL()
	// 如果请求中设置了过期时间，则按该时间设置新令牌的过期时间
	if expiredTTL > 0 {
		tk.SetExpiredAtByDuration(expiredTTL, 4) // 刷新令牌的过期时间是访问令牌的4倍
	}

	return tk, nil
}
