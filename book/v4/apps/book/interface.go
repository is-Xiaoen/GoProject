package book

import (
	"context"

	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/mcube/v2/types"
)

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

type DeleteBookRequest struct {
	DescribeBookRequest
}

type UpdateBookRequest struct {
	DescribeBookRequest
	CreateBookRequest
}

type DescribeBookRequest struct {
	Id uint
}

type BookSet struct {
	// 总共多少个
	Total int64 `json:"total"`
	// book清单
	Items []*Book `json:"items"`
}

func (b *BookSet) Add(item *Book) {
	b.Items = append(b.Items, item)
}

// type CommentSet struct {
// 	// 总共多少个
// 	Total int64 `json:"total"`
// 	// book清单
// 	Items []*Comment `json:"items"`
// }

// func (b *CommentSet) Add(item *Comment) {
// 	b.Items = append(b.Items, item)
// }

func NewCreateBookRequest() *CreateBookRequest {
	return (&CreateBookRequest{}).SetIsSale(false)
}

type CreateBookRequest struct {
	// type 用于要使用gorm 来自动创建和更新表的时候 才需要定义
	Title  string  `json:"title"  gorm:"column:title;type:varchar(200)" validate:"required"`
	Author string  `json:"author"  gorm:"column:author;type:varchar(200);index" validate:"required"`
	Price  float64 `json:"price"  gorm:"column:price" validate:"required"`
	// bool false
	// nil 是零值, false
	IsSale *bool `json:"is_sale"  gorm:"column:is_sale"`
}

// 这个请求对象的教育
func (r *CreateBookRequest) Validate() error {
	// validate := validator.New()
	// validate.Struct(r)
	return validator.Validate(r)
}

func (r *CreateBookRequest) SetIsSale(v bool) *CreateBookRequest {
	r.IsSale = &v
	return r
}

func NewQueryBookRequest() *QueryBookRequest {
	return &QueryBookRequest{
		// PageRequest{PageSize:20, PageNumber: 1}
		PageRequest: *request.NewDefaultPageRequest(),
	}
}

type QueryBookRequest struct {
	// PageSize   uint
	// PageNumber uint
	request.PageRequest
	// 关键字参数
	Keywords string `json:"keywords"`
}
