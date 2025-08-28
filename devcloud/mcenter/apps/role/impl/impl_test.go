package impl

import (
	"context"

	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/role"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/test"
)

var (
	impl role.Service
	ctx  = context.Background()
)

func init() {
	test.DevelopmentSetUp()
	impl = role.GetService()
}
