package impl_test

import (
	"context"

	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/policy"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/test"
)

var (
	impl policy.Service
	ctx  = context.Background()
)

func init() {
	test.DevelopmentSetUp()
	impl = policy.GetService()
}
