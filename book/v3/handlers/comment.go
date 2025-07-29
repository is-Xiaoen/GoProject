package handlers

import (
	"github.com/gin-gonic/gin"
)

var Comment = &CommentApiHandler{}

type CommentApiHandler struct {
}

func (h *CommentApiHandler) AddComment(ctx *gin.Context) {
	// Book.getBook()
	// Book.GetBook(id) -> BookInstance
	// controllers.Book.GetBook(ctx, controllers.NewGetBookRequest(ctx.Param("bn")))
}
