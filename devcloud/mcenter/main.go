package main

import (
	"github.com/infraboard/mcube/v2/ioc/server/cmd"

	// 加载的业务对象
	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps"

	// 非功能性模块
	_ "github.com/infraboard/mcube/v2/ioc/apps/apidoc/restful"
	_ "github.com/infraboard/mcube/v2/ioc/apps/health/restful"
	_ "github.com/infraboard/mcube/v2/ioc/apps/metric/restful"
)

func main() {
	// 启动
	cmd.Start()
}
