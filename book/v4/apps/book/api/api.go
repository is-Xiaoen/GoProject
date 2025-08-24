package api

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/is-Xiaoen/GoProject/book/v4/apps/book"

	// 引入Gin Root Router: *gin.Engine
	ioc_gin "github.com/infraboard/mcube/v2/ioc/config/gin"
	// 引入Gin Root Router: *gin.Engine
)

type BookApiHandler struct {
	ioc.ObjectImpl

	// 业务依赖
	svc book.Service
}

// 这个就是 API 的资源名称
// /api/book/v1/books
func (h *BookApiHandler) Name() string {
	return "books"
}

// 对象的初始化, 初始化对象的一些熟悉 &BookApiHandler{}
// 构造函数
// 当你这个对象初始化的时候，直接把的处理函数(ApiHandler注册给Server)
func (h *BookApiHandler) Init() error {
	h.svc = book.GetService()

	// 本地依赖
	// r := server.Gin
	// 框架托管, 通过容器获取 Server对象
	// 获取的 Gin Engine对象
	// ioc_gin.RootRouter()
	// URL 容器冲突,  book/comment
	// 怎么避免 2个业务API 不不冲突，加上业务板块的前缀，或者 对的名称
	// /<prefix>/<service_name>/<object_version>/<object_name>
	// http 接口前缀
	r := ioc_gin.ObjectRouter(h)

	// Book Restful API
	// List of books
	// /api/simple_api/v1/books
	r.GET("", h.queryBook)
	// Create new book
	// Body: HTTP Entity
	// /api/simple_api/v1/books
	r.POST("", h.createBook)
	return nil
}

func init() {
	ioc.Api().Registry(&BookApiHandler{})
}
