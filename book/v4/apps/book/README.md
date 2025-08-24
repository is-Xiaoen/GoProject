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


