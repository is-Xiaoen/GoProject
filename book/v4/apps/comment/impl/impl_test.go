package impl_test

import (
	"context"
	"os"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/is-Xiaoen/GoProject/book/v4/apps/comment/impl"
)

var ctx = context.Background()
var svc = impl.CommentServiceImpl{}

func init() {
	// import 后自动执行的逻辑
	// 工具对象的初始化, 需要的是绝对路径
	ioc.DevelopmentSetupWithPath(os.Getenv("workspaceFolder") + "/book/v4/application.toml")
}
