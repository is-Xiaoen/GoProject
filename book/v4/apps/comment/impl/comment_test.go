package impl_test

import (
	"testing"

	"github.com/is-Xiaoen/GoProject/book/v4/apps/comment"
)

func TestAddComment(t *testing.T) {
	ins, err := svc.AddComment(ctx, &comment.AddCommentRequest{
		BookId:  10,
		Comment: "评论测试",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
