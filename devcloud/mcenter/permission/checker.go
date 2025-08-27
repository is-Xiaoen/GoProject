package permission

import (
	"context"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/endpoint"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token"
	"github.com/rs/zerolog"
)

func init() {
	ioc.Config().Registry(&Checker{})
}

func Auth(v bool) (string, bool) {
	return endpoint.META_REQUIRED_AUTH_KEY, v
}

func Permission(v bool) (string, bool) {
	return endpoint.META_REQUIRED_PERM_KEY, v
}

func Resource(v string) (string, string) {
	return endpoint.META_RESOURCE_KEY, v
}

func Action(v string) (string, string) {
	return endpoint.META_ACTION_KEY, v
}

// 这个中间件也是对象, 认证与鉴权
// 通过路由装饰 来当中开关，控制怎么认证,是否开启认证，是否开启坚强，角色标记
type Checker struct {
	ioc.ObjectImpl
	log *zerolog.Logger

	token token.Service
	// policy policy.Service
}

// 中间件对象名称
func (c *Checker) Name() string {
	return "permission_checker"
}

// 对象初始化的优先级, 由于业务接口在Init函数里面 使用默认优先级0, 由大到小
// 框架是 899 898
// 框架的init函数调用完成，里面调用 这个对象的init函数, 实现了全局中间件
func (c *Checker) Priority() int {
	return gorestful.Priority() - 1
}

func (c *Checker) Init() error {
	c.log = log.Sub(c.Name())
	c.token = token.GetService()
	// c.policy = policy.GetService()

	// 注册认证中间件
	gorestful.RootRouter().Filter(c.Check)
	return nil
}

// 中间件的函数里面
func (c *Checker) Check(r *restful.Request, w *restful.Response, next *restful.FilterChain) {
	// 请求处理前, 对接口进行保护
	// 1. 知道用户当前访问的是哪个接口, 当前url 匹配到的路由是哪个
	// SelectedRoute, 它可以返回当前URL适配哪个路有， RouteReader
	// 封装了一个函数 来获取Meta信息 NewEntryFromRestRouteReader
	route := endpoint.NewEntryFromRestRouteReader(r.SelectedRoute())
	if route.RequiredAuth {
		// 校验身份
		tk, err := c.CheckToken(r)
		if err != nil {
			response.Failed(w, err)
			return
		}

		// 校验权限
		if err := c.CheckPolicy(r, tk, route); err != nil {
			response.Failed(w, err)
			return
		}
	}

	// 请求处理
	next.ProcessFilter(r, w)

	// 请求处理后
}

func (c *Checker) CheckToken(r *restful.Request) (*token.Token, error) {
	v := token.GetAccessTokenFromHTTP(r.Request)
	if v == "" {
		return nil, exception.NewUnauthorized("请先登录")
	}

	tk, err := c.token.ValiateToken(r.Request.Context(), token.NewValiateTokenRequest(v))
	if err != nil {
		return nil, err
	}

	// 如果校验成功，需要把 用户的身份信息，放到请求的上下文中，方便后面的逻辑获取
	// context.WithValue 来往ctx 添加 value
	// key: value, value token对象
	ctx := context.WithValue(r.Request.Context(), token.CTX_TOKEN_KEY, tk)

	// ctx 生成一个新的，继续往下传递
	r.Request = r.Request.WithContext(ctx)
	return tk, nil
}

func (c *Checker) CheckPolicy(r *restful.Request, tk *token.Token, route *endpoint.RouteEntry) error {
	return nil
}
