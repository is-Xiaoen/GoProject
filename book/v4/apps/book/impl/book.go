package impl

import (
	"context"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/types"
	"github.com/is-Xiaoen/GoProject/book/v3/config"
	"github.com/is-Xiaoen/GoProject/book/v4/apps/book"
)

// CreateBook implements book.Service.
func (b *BookServiceImpl) CreateBook(ctx context.Context, in *book.CreateBookRequest) (*book.Book, error) {
	// 自定义异常改造， 放到mcube
	// 自定义异常, exception 包, 统一放到一个公共库里面, mcube
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest("校验Book创建失败, %s", err)
	}

	bookInstance := &book.Book{CreateBookRequest: *in}

	// config对象改造
	// 数据入库(Grom), 补充自增Id的值
	if err := config.DB().Save(bookInstance).Error; err != nil {
		return nil, err
	}

	return bookInstance, nil
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
