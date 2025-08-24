package test

import (
	"os"

	"github.com/infraboard/mcube/v2/ioc"
	// 要注册哪些对象, Book, Comment

	_ "github.com/is-Xiaoen/GoProject/book/v4/apps/book/impl"
	_ "github.com/is-Xiaoen/GoProject/book/v4/apps/comment/impl"
)

func DevelopmentSet() {
	// import 后自动执行的逻辑
	// 工具对象的初始化, 需要的是绝对路径
	ioc.DevelopmentSetupWithPath(os.Getenv("workspaceFolder") + "/book/v4/application.toml")
}
