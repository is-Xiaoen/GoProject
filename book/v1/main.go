package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	// 对象Id
	ID uint `json:"id" gorm:"primaryKey;column:id"`

	// type 用于要使用gorm 来自动创建和更新表的时候 才需要定义
	Title  string  `json:"title"  gorm:"column:title;type:varchar(200)" validate:"required"`
	Author string  `json:"author"  gorm:"column:author;type:varchar(200);index" validate:"required"`
	Price  float64 `json:"price"  gorm:"column:price" validate:"required"`
	// bool false
	// nil 是零值, false
	IsSale *bool `json:"is_sale"  gorm:"column:is_sale"`
}

// TableName
func (b *Book) TableName() string {
	return "books"
}

// 初始化数据库
func setupDatabase() *gorm.DB {
	// 变量更新
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Book{}) // 自动迁移
	return db
}

var db = setupDatabase()

var h = &BookApiHandler{}

type BookApiHandler struct {
}

// 实现后端分页的
func (h *BookApiHandler) ListBook(ctx *gin.Context) {
	// /api/books?page_number=1&page_size=20
	pageNumber := ctx.Query("page_number")
	pn, err := strconv.ParseInt(pageNumber, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}
	pageSize := ctx.Query("page_size")
	ps, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	//
	bookList := []Book{}
	// select * from books
	// 通过sql的offset limte 来实现分页
	//  offset (page_number -1) * page_size, limit page_size
	// 2  offset 20, 20
	// 3  offset 40, 20
	// 4  offset 3 * 20, 20
	offset := (pn - 1) * ps
	if err := db.Offset(int(offset)).Limit(int(ps)).Find(&bookList).Error; err != nil {
		ctx.JSON(500, gin.H{"code": 500, "message": err.Error()})
		return
	}

	// 获取总数, 总共多少个, 总共有多少页
	ctx.JSON(200, bookList)

}

func (h *BookApiHandler) CreateBook(ctx *gin.Context) {
	// payload, err := io.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
	// 	return
	// }
	// defer ctx.Request.Body.Close()
	// // {"title": "Go语言"}

	// c.Request.Header.Get(key)
	// ctx.GetHeader("Authincation")

	// new(Book)
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

	// 数据入库(Grom), 补充自增Id的值
	if err := db.Save(bookInstance).Error; err != nil {
		ctx.JSON(400, gin.H{"code": 500, "message": err.Error()})
		return
	}

	// 返回响应
	ctx.JSON(200, bookInstance)
}

func (h *BookApiHandler) GetBook(ctx *gin.Context) {
	// URI
	bnStr := ctx.Param("bn")
	bn, err := strconv.ParseInt(bnStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}
	fmt.Println(bn)
}

func (h *BookApiHandler) UpdateBook(ctx *gin.Context) {
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
}

func (h *BookApiHandler) DeleteBook(ctx *gin.Context) {
	// URI
	bnStr := ctx.Param("bn")
	bn, err := strconv.ParseInt(bnStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}
	fmt.Println(bn)
}

func main() {
	server := gin.Default()

	// Book Restful API
	// List of books
	server.GET("/api/books", h.ListBook)
	// Create new book
	// Body: HTTP Entity
	server.POST("/api/books", h.CreateBook)
	// Get book by book number
	server.GET("/api/books/:bn", h.GetBook)
	// Update book
	server.PUT("/api/books/:bn", h.UpdateBook)
	// Delete book
	server.DELETE("/api/books/:bn", h.DeleteBook)

	if err := server.Run(":8080"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
