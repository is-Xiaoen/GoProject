package comment

import (
	"context"
)

const (
	APP_NAME = "comment"
)

// comment.Service, Comment的业务定义
type Service interface {
	// 为书籍添加评论
	AddComment(context.Context, *AddCommentRequest) (*Comment, error)
}

type AddCommentRequest struct {
	BookId  uint
	Comment string
}
