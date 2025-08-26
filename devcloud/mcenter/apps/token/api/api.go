package api

import (
	_ "embed"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
)

func init() {
	ioc.Api().Registry(&TokenRestulApiHandler{})
}

type TokenRestulApiHandler struct {
	ioc.ObjectImpl

	// 依赖控制器
	svc token.Service
}

func (h *TokenRestulApiHandler) Name() string {
	return token.APP_NAME
}

//go:embed docs/login.md
var loginApiDocNotes string

func (h *TokenRestulApiHandler) Init() error {
	h.svc = token.GetService()

	tags := []string{"用户登录"}
	ws := gorestful.ObjectRouter(h)
	ws.Route(ws.POST("").To(h.Login).
		Doc("颁发令牌(登录)").
		Notes(loginApiDocNotes).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(token.IssueTokenRequest{}).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}))

	ws.Route(ws.POST("/validate").To(h.ValiateToken).
		Doc("校验令牌").
		// Metadata(permission.Auth(true)).
		// Metadata(permission.Permission(false)).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(token.ValiateTokenRequest{}).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}))

	// ws.Route(ws.POST("/change_namespace").To(h.ChangeNamespce).
	// 	Doc("切换令牌访问空间").
	// 	// Metadata(permission.Auth(true)).
	// 	// Metadata(permission.Permission(false)).
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Reads(token.ChangeNamespceRequest{}).
	// 	Writes(token.Token{}).
	// 	Returns(200, "OK", token.Token{}))

	ws.Route(ws.DELETE("").To(h.Logout).
		Doc("撤销令牌(退出)").
		// Metadata(permission.Auth(true)).
		// Metadata(permission.Permission(false)).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(token.IssueTokenRequest{}).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}).
		Returns(404, "Not Found", nil))
	return nil
}
