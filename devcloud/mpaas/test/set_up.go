package test

import (
	"os"

	"github.com/infraboard/mcube/v2/ioc"
	// 要注册哪些对象, Book, Comment

	// 加载的业务对象
	_ "github.com/is-Xiaoen/GoProject/devcloud/mpaas/apps"
)

func DevelopmentSetUp() {
	// import 后自动执行的逻辑
	// 工具对象的初始化, 需要的是绝对路径
	ioc.DevelopmentSetupWithPath(os.Getenv("CONFIG_PATH"))
}
