package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Title string `json:"title"`
}

func main() {
	server := gin.Default()

	//Book Restful API
	//List of Books
	server.GET("/api/books/", func(ctx *gin.Context) {
		// /api/books?page_number=1&page_size=20
		//ctx.Query("page_number")
		//ctx.Query("page_size")
	})
	//Create new book
	//Create new Book
	server.POST("/api/books/", func(ctx *gin.Context) {
		// payload, err := io.ReadAll(ctx.Request.Body)
		// if err != nil {
		// 	ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		// 	return
		// }
		// defer ctx.Request.Body.Close()
		// // {"title": "Go语言"}

		bookInstance := &Book{}
		// // 通过JSON的 Struct Tag
		// // bookInstance.Title =  "Go语言"
		// if err := json.Unmarshal(payload, bookInstance); err != nil {
		// 	ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		// 	return
		// }
		// 获取到bookInstance
		if err := ctx.BindJSON(bookInstance); err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
			return
		}

		//数据入库 (Gorm)

		// 返回响应
		ctx.JSON(200, bookInstance)
	})

	//Get book by number
	server.GET("/api/books/:bn", func(ctx *gin.Context) {
		//URI
		bnStr := ctx.Param("bn")
		bn, err := strconv.ParseInt(bnStr, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
			return
		}
		fmt.Println(bn)
	})
	//Update book
	server.PUT("/api/books/:bn", func(ctx *gin.Context) {
		// URI
		bnStr := ctx.Param("bn")
		bn, err := strconv.ParseInt(bnStr, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
			return
		}
		fmt.Println(bn)

		// 读取body里面的参数
		bookInstance := &Book{}
		// 获取到bookInstance
		if err := ctx.BindJSON(bookInstance); err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
			return
		}
	})
	//Delete book
	server.DELETE("/api/books/:bn", func(ctx *gin.Context) {
		// URI
		bnStr := ctx.Param("bn")
		bn, err := strconv.ParseInt(bnStr, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
			return
		}
		fmt.Println(bn)
	})

	if err := server.Run(":8080"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
