package impl

import (
	"context"
	"github.com/is-Xiaoen/GoProject/book/v4/apps/book"
	"github.com/is-Xiaoen/GoProject/book/v4/apps/comment"
)

// AddComment implements comment.Service.
func (c *CommentServiceImpl) AddComment(ctx context.Context, in *comment.AddCommentRequest) (*comment.Comment, error) {
	// 能不能 直接Book Service的具体实现
	// (&impl.BookServiceImpl{}).DescribeBook(ctx, nil)
	// 依赖接口，面向接口编程, 不依赖具体实现
	book.GetService().DescribeBook(ctx, nil)
	panic("unimplemented")
}
