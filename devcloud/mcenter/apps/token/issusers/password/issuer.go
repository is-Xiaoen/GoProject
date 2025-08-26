package password

import (
	"context"
	"time"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/user"
)

func init() {
	// 将 PasswordTokenIssuer 实例注册到 ioc 配置中心，并设置默认过期时间
	ioc.Config().Registry(&PasswordTokenIssuer{
		ExpiredTTLSecond: 1 * 60 * 60, // 默认过期时间为 1 小时
	})
}

// PasswordTokenIssuer 实现了 token.Issuer 接口，用于通过用户名和密码颁发令牌
type PasswordTokenIssuer struct {
	ioc.ObjectImpl // 嵌入 ioc.ObjectImpl，提供基础的生命周期管理能力

	// 通过用户模块来判断用户凭证是否正确
	user user.Service

	// ExpiredTTLSecond Password 颁发的令牌过期时间，由系统配置，不允许用户自己设置
	ExpiredTTLSecond int `json:"expired_ttl_second" toml:"expired_ttl_second" yaml:"expired_ttl_second" env:"EXPIRED_TTL_SECOND"`

	// expiredDuration 是根据 ExpiredTTLSecond 计算出的时间持续时长
	expiredDuration time.Duration
}

// Name 返回颁发器的名称，用于在 ioc 容器中注册
func (p *PasswordTokenIssuer) Name() string {
	return "password_token_issuer"
}

// Init 是 ioc 容器在服务启动时调用的初始化方法
func (p *PasswordTokenIssuer) Init() error {
	// 从 ioc 容器中获取用户服务实例
	p.user = user.GetService()
	// 根据配置的秒数计算过期时间
	p.expiredDuration = time.Duration(p.ExpiredTTLSecond) * time.Second
	// 将自身注册为 token 包中的颁发器，并指定名称
	token.RegistryIssuer(token.ISSUER_PASSWORD, p)
	return nil
}

// IssueToken 实现了 token.Issuer 接口的 IssueToken 方法
func (p *PasswordTokenIssuer) IssueToken(ctx context.Context, parameter token.IssueParameter) (*token.Token, error) {
	// 1. 查询用户
	uReq := user.NewDescribeUserRequestByUserName(parameter.Username())
	u, err := p.user.DescribeUser(ctx, uReq)
	if err != nil {
		// 如果用户不存在，则返回“未授权”错误
		if exception.IsNotFoundError(err) {
			return nil, exception.NewUnauthorized("%s", err)
		}
		// 返回其他类型的错误
		return nil, err
	}

	// 2. 比对密码
	err = u.CheckPassword(parameter.Password())
	if err != nil {
		// 如果密码不匹配，返回错误
		return nil, err
	}

	// 3. 颁发令牌
	tk := token.NewToken()
	// 设置令牌关联的用户信息
	tk.UserId = u.Id
	tk.UserName = u.UserName
	tk.IsAdmin = u.IsAdmin
	// 设置令牌的过期时间，刷新令牌的过期时间是访问令牌的4倍
	tk.SetExpiredAtByDuration(p.expiredDuration, 4)

	return tk, nil
}
