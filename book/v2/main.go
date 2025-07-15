package main

import (
	"fmt"
	"github.com/is-Xiaoen/GoProject/book/v2/config"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Book 结构体定义
type Book struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// 初始化数据库
func setupDatabase() *gorm.DB {
	mc := config.C().MySQL
	// dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mc.Username,
		mc.Password,
		mc.Host,
		mc.Port,
		mc.DB,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Book{}) // 自动迁移
	return db
}

func main() {
	// 加载配置
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "application.yaml"
	}
	config.LoadConfigFromYaml(path)

	r := gin.Default()
	db := setupDatabase()

	// 创建书籍
	r.POST("/books", func(c *gin.Context) {
		var book Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&book)
		c.JSON(http.StatusCreated, book)
	})

	// 获取所有书籍
	r.GET("/books", func(c *gin.Context) {
		var books []Book
		db.Find(&books)
		c.JSON(http.StatusOK, books)
	})

	// 根据 ID 获取书籍
	r.GET("/books/:id", func(c *gin.Context) {
		var book Book
		id := c.Param("id")
		if err := db.First(&book, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusOK, book)
	})

	// 更新书籍
	r.PUT("/books/:id", func(c *gin.Context) {
		var book Book
		id := c.Param("id")
		if err := db.First(&book, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Save(&book)
		c.JSON(http.StatusOK, book)
	})

	// 删除书籍
	r.DELETE("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&Book{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	})

	ac := config.C().Application
	r.Run(fmt.Sprintf("%s:%d", ac.Host, ac.Port)) // 启动服务
}
