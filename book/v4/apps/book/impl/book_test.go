package impl_test

import (
	"testing"

	"github.com/is-Xiaoen/GoProject/book/v4/apps/book"
)

//

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
