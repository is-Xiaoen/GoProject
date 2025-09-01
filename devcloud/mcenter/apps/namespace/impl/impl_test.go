package impl_test

import (
	"context"

	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/namespace"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/test"
)

var (
	impl namespace.Service
	ctx  = context.Background()
)

func init() {
	test.DevelopmentSetUp()
	impl = namespace.GetService()
}
