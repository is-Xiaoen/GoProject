package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/is-Xiaoen/GoProject/book/v3/config"
	"github.com/is-Xiaoen/GoProject/book/v3/controllers"
	"github.com/is-Xiaoen/GoProject/book/v3/models"
)

var Book = &BookApiHandler{}

type BookApiHandler struct {
}

func (h *BookApiHandler) Registry(r gin.IRouter) {
	// Book Restful API
	// List of books
	r.GET("/api/books", h.listBook)
	// Create new book
	// Body: HTTP Entity
	r.POST("/api/books", h.createBook)
	// Get book by book number
	r.GET("/api/books/:bn", h.getBook)
	// Update book
	r.PUT("/api/books/:bn", h.updateBook)
	// Delete book
	r.DELETE("/api/books/:bn", h.deleteBook)
}

// 实现后端分页的
func (h *BookApiHandler) listBook(ctx *gin.Context) {
	set := &models.BookSet{}

	// List<*Book>
	//  *Set[T]
	// types.New[*Book]()

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

	query := config.DB().Model(&models.Book{})
	// 关键字过滤
	kws := ctx.Query("keywords")
	if kws != "" {
		// where title like %kws%
		query = query.Where("title LIKE ?", "%"+kws+"%")
	}

	// 其他过滤条件

	// select * from books
	// 通过sql的offset limte 来实现分页
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

func (h *BookApiHandler) createBook(ctx *gin.Context) {
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
	bookSpecInstance := &models.BookSpec{}
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

	book, err := controllers.Book.CreateBook(ctx.Request.Context(), bookSpecInstance)
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	// 返回响应
	ctx.JSON(http.StatusCreated, book)
}

func (h *BookApiHandler) getBook(ctx *gin.Context) {
	bnInt, err := strconv.ParseInt(ctx.Param("bn"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	book, err := controllers.Book.GetBook(ctx, controllers.NewGetBookRequest(int(bnInt)))
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	ctx.JSON(200, book)
}

func (h *BookApiHandler) updateBook(ctx *gin.Context) {
	bnStr := ctx.Param("bn")
	bn, err := strconv.ParseInt(bnStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	// 读取body里面的参数
	bookInstance := &models.Book{
		Id: uint(bn),
	}
	// 获取到bookInstance
	if err := ctx.BindJSON(&bookInstance.BookSpec); err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	if err := config.DB().Where("id = ?", bookInstance.Id).Updates(bookInstance).Error; err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}

	ctx.JSON(200, bookInstance)
}

func (h *BookApiHandler) deleteBook(ctx *gin.Context) {
	if err := config.DB().Where("id = ?", ctx.Param("bn")).Delete(&models.Book{}).Error; err != nil {
		ctx.JSON(400, gin.H{"code": 400, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, "ok")
}
