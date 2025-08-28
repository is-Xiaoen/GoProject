package test

import (
	"fmt"
	"os"

	"github.com/infraboard/mcube/v2/ioc"

	// 加载的业务对象
	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps"
)

func DevelopmentSetUp() {
	// import 后自动执行的逻辑
	// 工具对象的初始化, 需要的是绝对路径
	fmt.Println(os.Getenv("CONFIG_PATH"))
	ioc.DevelopmentSetupWithPath(os.Getenv("CONFIG_PATH"))
}
