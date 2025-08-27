package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/user"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/permission"
)

func init() {
	ioc.Api().Registry(&UserRestfulApiHandler{})
}

type UserRestfulApiHandler struct {
	ioc.ObjectImpl

	// 依赖控制器
	svc user.Service
}

func (h *UserRestfulApiHandler) Name() string {
	return "users"
}

func (h *UserRestfulApiHandler) Init() error {
	h.svc = user.GetService()

	tags := []string{"用户登录"}
	ws := gorestful.ObjectRouter(h)
	// required_auth=true/false
	ws.Route(ws.GET("").To(h.QueryUser).
		Doc("用户列表查询").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 这个开关怎么生效
		// 中间件需求读取接口的描述信息，来决定是否需要认证
		Metadata(permission.Auth(true)).
		Param(restful.QueryParameter("page_size", "分页大小").DataType("integer")).
		Param(restful.QueryParameter("page_number", "页码").DataType("integer")).
		Writes(Set{}).
		Returns(200, "OK", Set{}))

	return nil
}

// *types.Set[*User]
// 返回的泛型, API Doc这个工具 不支持泛型
type Set struct {
	Total int64       `json:"total"`
	Items []user.User `json:"items"`
}
