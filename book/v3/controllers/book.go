package controllers

import (
	"context"

	"github.com/is-Xiaoen/GoProject/book/v3/config"
	"github.com/is-Xiaoen/GoProject/book/v3/exception"
	"github.com/is-Xiaoen/GoProject/book/v3/models"
	"gorm.io/gorm"
)

var Book = &BookController{}

type BookController struct {
}

func NewGetBookRequest(bookNumber int) *GetBookRequest {
	return &GetBookRequest{
		BookNumber: bookNumber,
	}
}

type GetBookRequest struct {
	BookNumber int
	// RequestId  string
	// ...
}

// GetBook 核心功能
// ctx: Trace, 支持请求的取消, request_id
// GetBookRequest 为什么要把他封装为1个对象, GetBook(ctx context.Context, BookNumber string), 保证你的接口的签名的兼容性
// BookController.GetBook(, "")
func (c *BookController) GetBook(ctx context.Context, in *GetBookRequest) (*models.Book, error) {
	// context.WithValue(ctx, "request_id", 111)
	// ctx.Value("request_id")

	config.L().Debug().Msgf("get book: %d", in.BookNumber)

	bookInstance := &models.Book{}
	// 需要从数据库中获取一个对象
	if err := config.DB().Where("id = ?", in.BookNumber).Take(bookInstance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.ErrNotFound("book number: %d not found", in.BookNumber)
		}
		return nil, err
	}

	return bookInstance, nil
}

func (c *BookController) CreateBook(ctx context.Context, in *models.BookSpec) (*models.Book, error) {
	// 有没有能够检查某个字段是否是必须填
	// Gin 集成 validator这个库, 通过 struct tag validate 来表示这个字段是否允许为空
	// validate:"required"
	// 在数据Bind的时候，这个逻辑会自动运行
	// if bookSpecInstance.Author == "" {
	// 	ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
	// 	return
	// }

	bookInstance := &models.Book{BookSpec: *in}

	// 数据入库(Grom), 补充自增Id的值
	if err := config.DB().Save(bookInstance).Error; err != nil {
		return nil, err
	}

	return bookInstance, nil
}

func (c *BookController) UpdateBook() {
	// update(obj)
	// config.DB().Updates()
}

func (c *BookController) update(ctx context.Context, obj models.Book) error {
	// obj.UpdateTime = now()
	// config.DB().Updates()
	return nil
}
