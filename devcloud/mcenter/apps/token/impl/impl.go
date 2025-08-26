package impl

import (
	"time"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/user"
	"github.com/rs/zerolog"
)

// init 函数是 Go 语言的特殊函数，在包被导入时自动执行
func init() {
	// 将 TokenServiceImpl 实例注册到 ioc 容器中
	ioc.Controller().Registry(&TokenServiceImpl{
		// 自动刷新配置默认为 true
		AutoRefresh: true,
		// 刷新 TTL（生存时间）默认为 1 小时
		RereshTTLSecond: 1 * 60 * 60,
	})
}

// 确保 TokenServiceImpl 类型实现了 token.Service 接口
var _ token.Service = (*TokenServiceImpl)(nil)

// TokenServiceImpl 是 token.Service 接口的具体实现
type TokenServiceImpl struct {
	ioc.ObjectImpl // 嵌入 ioc.ObjectImpl，提供基础的生命周期管理能力
	user           user.Service
	log            *zerolog.Logger
	// policy policy.PermissionService // 权限服务，目前被注释掉

	// AutoRefresh 控制是否自动刷新令牌的过期时间，而不是生成一个新令牌
	AutoRefresh bool `json:"auto_refresh" toml:"auto_refresh" yaml:"auto_refresh" env:"AUTO_REFRESH"`
	// RereshTTLSecond 刷新令牌的生存时间（秒）
	RereshTTLSecond uint64 `json:"refresh_ttl" toml:"refresh_ttl" yaml:"refresh_ttl" env:"REFRESH_TTL"`
	// MaxActiveApiToken 单个用户最多可以拥有的活跃 API 令牌数量，用于安全限制
	MaxActiveApiToken uint8 `json:"max_active_api_token" toml:"max_active_api_token" yaml:"max_active_api_token" env:"MAX_ACTIVE_API_TOKEN"`

	// refreshDuration 是根据 RereshTTLSecond 计算出的时间持续时长
	refreshDuration time.Duration
}

// Init 是 ioc 容器在服务启动时调用的初始化方法
func (i *TokenServiceImpl) Init() error {
	// 初始化子日志记录器
	i.log = log.Sub(i.Name())
	// 从 ioc 容器获取用户服务实例
	i.user = user.GetService()
	// i.policy = policy.GetService()

	// 根据配置的秒数计算刷新持续时间
	i.refreshDuration = time.Duration(i.RereshTTLSecond) * time.Second

	// 如果配置了自动迁移数据库表结构
	if datasource.Get().AutoMigrate {
		// 自动迁移 token.Token 结构体对应的数据库表
		err := datasource.DB().AutoMigrate(&token.Token{})
		if err != nil {
			return err
		}
	}
	return nil
}

// Name 返回服务的名称，用于在 ioc 容器中注册和查找
func (i *TokenServiceImpl) Name() string {
	return token.APP_NAME
}
