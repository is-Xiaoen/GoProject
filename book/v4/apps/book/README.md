# Book 业务分区

## 定义Book业务逻辑

业务功能: CRUD
1. 创建书籍(录入)
2. Book列表查询
3. Book详情查询
4. Book更新
5. Book删除

通过Go语言的里面的接口 来定义描述业务功能

```go
// book.Service, Book的业务定义
type Service interface {
// 1. 创建书籍(录入)
CreateBook(context.Context, *CreateBookRequest) (*Book, error)
// 2. Book列表查询
QueryBook(context.Context, *QueryBookRequest) (*types.Set[*Book], error)
// 3. Book详情查询
DescribeBook(context.Context, *DescribeBookRequest) (*Book, error)
// 4. Book更新
UpdateBook(context.Context, *UpdateBookRequest) (*Book, error)
// 5. Book删除
DeleteBook(context.Context, *DeleteBookRequest) error
}
```

## 业务的具体实现(TDD： Test Drive Develop)

1. BookServiceImpl

```go
// CreateBook implements book.Service.
func (b *BookServiceImpl) CreateBook(context.Context, *book.CreateBookRequest) (*book.Book, error) {
	panic("unimplemented")
}

// DeleteBook implements book.Service.
func (b *BookServiceImpl) DeleteBook(context.Context, *book.DeleteBookRequest) error {
	panic("unimplemented")
}

// DescribeBook implements book.Service.
func (b *BookServiceImpl) DescribeBook(context.Context, *book.DescribeBookRequest) (*book.Book, error) {
	panic("unimplemented")
}

// QueryBook implements book.Service.
func (b *BookServiceImpl) QueryBook(context.Context, *book.QueryBookRequest) (*types.Set[*book.Book], error) {
	panic("unimplemented")
}

// UpdateBook implements book.Service.
func (b *BookServiceImpl) UpdateBook(context.Context, *book.UpdateBookRequest) (*book.Book, error) {
	panic("unimplemented")
}
```

## 编写单元测试

```go
func TestCreateBook(t *testing.T) {
	req := book.NewCreateBookRequest()
	req.SetIsSale(true)
	req.Title = "Go语言V4"
	req.Author = "will"
	req.Price = 10
	ins, err := svc.CreateBook(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
```

## 业务对象注册(ioc controller)

手动维护
```sh
pkg gloab

bookContrller = xxx
commentContrller = xx
...
```

通过容器来维护对象

```go
// Book业务的具体实现
type BookServiceImpl struct {
	ioc.ObjectImpl
}

// 返回对象的名称, 因此我需要 服务名称
// 当前的MySQLBookServiceImpl 是 service book.APP_NAME 的 一个具体实现
// 当前的MongoDBBookServiceImpl 是 service book.APP_NAME 的 一个具体实现
func (s *BookServiceImpl) Name() string {
	return book.APP_NAME
}

func init() {
	ioc.Controller().Registry(&BookServiceImpl{})
}
```

## 面向接口

对象取处理, 断言他满足业务接口，然后我们以接口的方式来使用

```go
func GetService() Service {
	return ioc.Controller().Get(APP_NAME).(Service)
}

const (
	APP_NAME = "book"
)
```

第三方模块，可以依赖 接口进行开发
```go
// AddComment implements comment.Service.
func (c *CommentServiceImpl) AddComment(ctx context.Context, in *comment.AddCommentRequest) (*comment.Comment, error) {
	// 能不能 直接Book Service的具体实现
	// (&impl.BookServiceImpl{}).DescribeBook(ctx, nil)
	// 依赖接口，面向接口编程, 不依赖具体实现
	book.GetService().DescribeBook(ctx, nil)
	panic("unimplemented")
}
```

## 开发API

接口是需求, 对业务进行设计, 可以选择把这些能力 以那种接口的访问对外提供服务

1. 不对外提供接口，仅仅作为其他的业务的依赖
2. (API)对外提供 HTTP接口, RESTful接口
3. (API)对内提供 RPC接口(JSON RPC/GRPC/thrift)

1. 开发业务功能
```go
func (h *BookApiHandler) createBook(ctx *gin.Context) {
	req := book.NewCreateBookRequest()
	if err := ctx.BindJSON(req); err != nil {
		response.Failed(ctx, err)
		return
	}

	ins, err := h.svc.CreateBook(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	// 返回响应
	response.Success(ctx, ins)
}
```

2. 注册路由
```go
type BookApiHandler struct {
	ioc.ObjectImpl

	// 业务依赖
	svc book.Service
}

func (h *BookApiHandler) Name() string {
	return "books"
}

// 对象的初始化, 初始化对象的一些熟悉 &BookApiHandler{}
// 构造函数
// 当你这个对象初始化的时候，直接把的处理函数(ApiHandler注册给Server)
func (h *BookApiHandler) Init() error {
	h.svc = book.GetService()
	r := ioc_gin.ObjectRouter(h)

	r.GET("", h.queryBook)
	r.POST("", h.createBook)
	return nil
}

func init() {
	ioc.Api().Registry(&BookApiHandler{})
}
```

## 业务注册

每写完一个业务，就需要在 注册到ioc(注册表)
```go
// 业务加载区, 选择性的价值的业务处理对象

import (
	// Api Impl
	_ "122.51.31.227/go-course/go18/book/v4/apps/book/api"

	// Service Impl
	_ "122.51.31.227/go-course/go18/book/v4/apps/book/impl"
	_ "122.51.31.227/go-course/go18/book/v4/apps/comment/impl"
)
```


## 启动服务

```go
import (
	"github.com/infraboard/mcube/v2/ioc/server/cmd"

	// 业务对象
	_ "122.51.31.227/go-course/go18/book/v4/apps"
)

func main() {
	// ioc框架 加载对象, 注入对象, 配置对象
	// server.Gin.Run()
	// application.Get().AppName
	// http.Get().Host
	// server.DefaultConfig.ConfigFile.Enabled = true
	// server.DefaultConfig.ConfigFile.Path = "application.toml"
	// server.Run(context.Background())
	// 不能指定配置文件逻辑
	// 使用者来说，体验不佳

	// ioc 直接提供server, 直接run就行了，
	// mcube 包含 一个 gin Engine
	// CLI, start 指令 -f 指定配置文件
	cmd.Start()
}
```

```toml
[app]
  name = "simple_api"
  description = "app desc"
  address = "localhost"
  encrypt_key = "defualt app encrypt key"

[datasource]
  provider = "mysql"
  host = "127.0.0.1"
  port = 3306
  database = "go18"
  username = "root"
  password = "123456"
  auto_migrate = false
  debug = false

[http]
  host = "127.0.0.1"
  port = 8010
  path_prefix = "api"

[comment]
  max_comment_per_book = 200
```

```sh
➜  v4 git:(main) ✗ go run main.go -f application.toml start
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

2025-05-25T15:59:33+08:00 INFO   config/gin/framework.go:41 > enable gin recovery component:GIN_WEBFRAMEWORK
200
[GIN-debug] GET    /api/simple_api/v1/books  --> 122.51.31.227/go-course/go18/book/v4/apps/book/api.(*BookApiHandler).queryBook-fm (4 handlers)
[GIN-debug] POST   /api/simple_api/v1/books  --> 122.51.31.227/go-course/go18/book/v4/apps/book/api.(*BookApiHandler).createBook-fm (4 handlers)
2025-05-25T15:59:33+08:00 INFO   ioc/server/server.go:74 > loaded configs: [app.v1 trace.v1 log.v1 validator.v1 gin_webframework.v1 datasource.v1 grpc.v1 http.v1] component:SERVER
2025-05-25T15:59:33+08:00 INFO   ioc/server/server.go:75 > loaded controllers: [comment.v1 book.v1] component:SERVER
2025-05-25T15:59:33+08:00 INFO   ioc/server/server.go:76 > loaded apis: [books.v1] component:SERVER
2025-05-25T15:59:33+08:00 INFO   ioc/server/server.go:77 > loaded defaults: [] component:SERVER
2025-05-25T15:59:33+08:00 INFO   config/http/http.go:144 > HTTP服务启动成功, 监听地址: 127.0.0.1:8010 component:HTTP
```

## 总结

业务分区框架, 我们专注于业务对象的开发, mcube相对于一个工具箱，承接其他非业务的公共功能


## 其他非功能需求

工具箱 提供很多工具，开箱即用, 比如health check, 比如metrics

```go
	// 健康检查
	_ "github.com/infraboard/mcube/v2/ioc/apps/health/gin"
	// metrics
	_ "github.com/infraboard/mcube/v2/ioc/apps/metric/gin"
```


[metric.v1 books.v1 health.v1] metric, health 使用注入的对象
```sh
➜  v4 git:(main) ✗ go run main.go -f application.toml start
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

2025-05-25T16:06:42+08:00 INFO   config/gin/framework.go:41 > enable gin recovery component:GIN_WEBFRAMEWORK
200
[GIN-debug] GET    /metrics/                 --> github.com/infraboard/mcube/v2/ioc/apps/metric/gin.(*ginHandler).Registry.func1 (5 handlers)
2025-05-25T16:06:42+08:00 INFO   metric/gin/metric.go:89 > Get the Metric using http://127.0.0.1:8010/metrics component:METRIC
[GIN-debug] GET    /api/simple_api/v1/books  --> 122.51.31.227/go-course/go18/book/v4/apps/book/api.(*BookApiHandler).queryBook-fm (5 handlers)
[GIN-debug] POST   /api/simple_api/v1/books  --> 122.51.31.227/go-course/go18/book/v4/apps/book/api.(*BookApiHandler).createBook-fm (5 handlers)
[GIN-debug] GET    /healthz/                 --> github.com/infraboard/mcube/v2/ioc/apps/health/gin.(*HealthChecker).HealthHandleFunc-fm (5 handlers)
2025-05-25T16:06:42+08:00 INFO   health/gin/check.go:55 > Get the Health using http://127.0.0.1:8010/healthz component:HEALTH_CHECK
2025-05-25T16:06:42+08:00 INFO   ioc/server/server.go:74 > loaded configs: [app.v1 trace.v1 log.v1 validator.v1 gin_webframework.v1 datasource.v1 grpc.v1 http.v1] component:SERVER
2025-05-25T16:06:42+08:00 INFO   ioc/server/server.go:75 > loaded controllers: [comment.v1 book.v1] component:SERVER
2025-05-25T16:06:42+08:00 INFO   ioc/server/server.go:76 > loaded apis: [metric.v1 books.v1 health.v1] component:SERVER
2025-05-25T16:06:42+08:00 INFO   ioc/server/server.go:77 > loaded defaults: [] component:SERVER
2025-05-25T16:06:42+08:00 INFO   config/http/http.go:144 > HTTP服务启动成功, 监听地址: 127.0.0.1:8010 component:HTTP
```