package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BookSet struct {
	//总共多少个
	Total int64 `json:"total"`
	//book清单
	Items []*Book `json:"items"`
}

type Book struct {
	// 对象Id
	Id uint `json:"id" gorm:"primaryKey;column:id"`

	BookSpec
}

type BookSpec struct {
	// type 用于要使用gorm 来自动创建和更新表的时候 才需要定义
	Title  string  `json:"title"  gorm:"column:title;type:varchar(200)" validate:"required"`
	Author string  `json:"author"  gorm:"column:author;type:varchar(200);index" validate:"required"`
	Price  float64 `json:"price"  gorm:"column:price" validate:"required"`
	// bool false
	// nil 是零值, false
	IsSale *bool `json:"is_sale"  gorm:"column:is_sale"`
}

// books
func (b *Book) TableName() string {
	return "books"
}

// setupDatabase 初始化数据库
func setupDatabase() *gorm.DB {
	// 变量更新
	dsn := "root:123456@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
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
	set := &BookSet{}

	// 给默认值
	pn, ps := 1, 20
	// /api/books?page_number=1&page_size=20
	pageNumber := ctx.Query("page_number")
	if pageNumber != "" {
		pnInt, err := strconv.ParseInt(pageNumber, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
			return
		}
		pn = int(pnInt)
	}

	pageSize := ctx.Query("page_size")
	if pageSize != "" {
		psInt, err := strconv.ParseInt(pageSize, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
			return
		}
		ps = int(psInt)
	}

	query := db.Model(&Book{})
	//关键字过滤
	kws := ctx.Query("keywords")
	if kws != "" {
		query = query.Where("title like ?", "%"+kws+"%")
	}

	// 其他过滤条件

	// select * from books
	// 通过sql的offset limite 来实现分页
	//  offset (page_number -1) * page_size, limit page_size
	// 2  offset 20, 20
	// 3  offset 40, 20
	// 4  offset 3 * 20, 20
	offset := (pn - 1) * ps
	if err := query.Count(&set.Total).Offset(int(offset)).Limit(int(ps)).Find(&set.Items).Error; err != nil {
		ctx.JSON(500, gin.H{"code": 500, "message": err.Error()})
		return
	}

	// 获取总数, 总共多少个, 总共有多少页
	ctx.JSON(200, set)
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
	bookSpecInstance := &BookSpec{}
	// // 通过JSON的 Struct Tag
	// // bookInstance.Title =  "Go语言"
	// if err := json.Unmarshal(payload, bookInstance); err != nil {
	// 	ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
	// 	return
	// }
	// 获取到bookInstance
	// 参数是不是为空
	if err := ctx.BindJSON(bookSpecInstance); err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	// 有没有能够检查某个字段是否是必须填
	// Gin 集成 validator这个库, 通过 struct tag validate 来表示这个字段是否允许为空
	// validate:"required"
	// 在数据Bind的时候，这个逻辑会自动运行
	// if bookSpecInstance.Author == "" {
	// 	ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
	// 	return
	// }

	bookInstance := &Book{BookSpec: *bookSpecInstance}

	// 数据入库(Grom), 补充自增Id的值
	if err := db.Save(bookInstance).Error; err != nil {
		ctx.JSON(400, gin.H{"code": 500, "message": err.Error()})
		return
	}

	// 返回响应
	ctx.JSON(http.StatusCreated, bookInstance)
}

func (h *BookApiHandler) GetBook(ctx *gin.Context) {
	bookInstance := &Book{}
	//需要从数据库中获取一个对象
	if err := db.Where("id = ?", ctx.Param("bn")).Take(bookInstance).Error; err != nil {
		ctx.JSON(400, gin.H{"code": 500, "message": err.Error()})
		return
	}

	ctx.JSON(200, bookInstance)
}

func (h *BookApiHandler) UpdateBook(ctx *gin.Context) {
	bnStr := ctx.Param("bn")
	bn, err := strconv.ParseInt(bnStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}
	fmt.Println(bn)

	// 读取body里面的参数
	bookInstance := &Book{
		Id: uint(bn),
	}
	// 获取到bookInstance
	if err := ctx.BindJSON(&bookInstance.BookSpec); err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	if err := db.Where("id = ?", bookInstance.Id).Updates(bookInstance).Error; err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}
	ctx.JSON(200, bookInstance)
}

func (h *BookApiHandler) DeleteBook(ctx *gin.Context) {
	if err := db.Where("id = ?", ctx.Param("bn")).Delete(&Book{}).Error; err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, "ok")
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
