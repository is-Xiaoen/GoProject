package impl_test

import (
	"context"

	"github.com/is-Xiaoen/GoProject/devcloud/mpaas/apps/application"
	"github.com/is-Xiaoen/GoProject/devcloud/mpaas/test"
)

var (
	svc application.Service
	ctx = context.Background()
)

func init() {
	test.DevelopmentSetUp()
	svc = application.GetService()
}
