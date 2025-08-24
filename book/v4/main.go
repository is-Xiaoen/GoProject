package main

import (
	"github.com/infraboard/mcube/v2/ioc/server/cmd"

	// 业务对象
	_ "github.com/is-Xiaoen/GoProject/book/v4/apps"

	// 健康检查
	_ "github.com/infraboard/mcube/v2/ioc/apps/health/gin"
	// metrics
	_ "github.com/infraboard/mcube/v2/ioc/apps/metric/gin"
)

func main() {
	// ioc框架 加载对象, 注入对象, 配置对象
	// server.Gin.Run()
	// application.Get().AppName
	// http.Get().Host
	// server.DefaultConfig.ConfigFile.Enabled = true
	// server.DefaultConfig.ConfigFile.Path = "application.toml"
	// server.Run(context.Background())
	// 不能指定配置文件逻辑
	// 使用者来说，体验不佳

	// ioc 直接提供server, 直接run就行了，
	// mcube 包含 一个 gin Engine
	// CLI, start 指令 -f 指定配置文件
	cmd.Start()
}
