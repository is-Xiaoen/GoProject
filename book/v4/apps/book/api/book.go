package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/http/gin/response"
	"github.com/is-Xiaoen/GoProject/book/v4/apps/book"
)

func (h *BookApiHandler) queryBook(ctx *gin.Context) {
	// 给默认值
	req := book.NewQueryBookRequest()
	req.Keywords = ctx.Query("keywords")
	// /api/books?page_number=1&page_size=20
	pageNumber := ctx.Query("page_number")
	if pageNumber != "" {
		pnInt, err := strconv.ParseInt(pageNumber, 10, 64)
		if err != nil {
			response.Failed(ctx, err)
			return
		}
		req.PageNumber = uint64(pnInt)
	}

	pageSize := ctx.Query("page_size")
	if pageSize != "" {
		psInt, err := strconv.ParseInt(pageSize, 10, 64)
		if err != nil {
			response.Failed(ctx, err)
			return
		}
		req.PageSize = uint64(psInt)
	}

	set, err := h.svc.QueryBook(ctx.Request.Context(), req)
	if err != nil {
		// 针对Response的统一封装, 已经落到 mcube
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, set)
}

func (h *BookApiHandler) createBook(ctx *gin.Context) {
	req := book.NewCreateBookRequest()
	if err := ctx.BindJSON(req); err != nil {
		response.Failed(ctx, err)
		return
	}

	ins, err := h.svc.CreateBook(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	// 返回响应
	response.Success(ctx, ins)
}
