package impl

import (
	"context"

	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/user"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/test"
)

var (
	impl user.Service
	ctx  = context.Background()
)

func init() {
	test.DevelopmentSet()
	impl = user.GetService()
}
