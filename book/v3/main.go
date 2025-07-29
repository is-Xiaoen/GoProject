package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/is-Xiaoen/GoProject/book/v3/config"
	"github.com/is-Xiaoen/GoProject/book/v3/handlers"
)

func main() {
	// 加载配置
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "application.yaml"
	}
	config.LoadConfigFromYaml(path)

	server := gin.Default()

	handlers.Book.Registry(server)

	ac := config.C().Application
	// 启动服务
	if err := server.Run(fmt.Sprintf("%s:%d", ac.Host, ac.Port)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
