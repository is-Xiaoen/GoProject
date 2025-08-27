package api

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/user"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
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
	ws.Route(ws.GET("").To(h.QueryUser).
		Doc("用户列表查询").
		Metadata(restfulspec.KeyOpenAPITags, tags).
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
