package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	server := gin.Default()
	server.GET("/hello/", func(c *gin.Context) {
		c.String(200, "Gin Hello World!")
	})
	if err := server.Run(":8080"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
