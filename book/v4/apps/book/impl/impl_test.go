package impl_test

import (
	"context"

	"github.com/is-Xiaoen/GoProject/book/v4/apps/book"
	"github.com/is-Xiaoen/GoProject/book/v4/test"
)

var ctx = context.Background()
var svc book.Service

func init() {
	test.DevelopmentSet()

	svc = book.GetService()
}
