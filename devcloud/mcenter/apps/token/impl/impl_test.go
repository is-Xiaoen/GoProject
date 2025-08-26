package impl_test

import (
	"context"

	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/test"
)

var (
	svc token.Service
	ctx = context.Background()
)

func init() {
	test.DevelopmentSet()
	svc = token.GetService()
}
