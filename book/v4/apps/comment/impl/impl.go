package impl

import (
	"fmt"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/is-Xiaoen/GoProject/book/v4/apps/comment"
)

func init() {
	ioc.Controller().Registry(&CommentServiceImpl{
		MaxCommentPerBook: 100,
	})
}

// 怎么知道他有没有实现该业务, 可以通过类型约束
// var _ book.Service = &BookServiceImpl{}

//	&BookServiceImpl 的 nil对象
//
// int64(1)  int64 1
// *BookServiceImpl(nil)
var _ comment.Service = (*CommentServiceImpl)(nil)

// Book业务的具体实现
type CommentServiceImpl struct {
	ioc.ObjectImpl

	// Comment最大限制
	MaxCommentPerBook int `toml:"max_comment_per_book"`
}

func (s *CommentServiceImpl) Init() error {
	// 当前对象，已经读取了配置文件
	fmt.Println(s.MaxCommentPerBook)
	return nil
}

func (s *CommentServiceImpl) Name() string {
	return comment.APP_NAME
}
