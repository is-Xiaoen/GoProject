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

编写单元测试

