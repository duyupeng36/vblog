package api

import (
	"net/http"
	"strconv"
	"vblog/apps/blog"
	"vblog/utils"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateBlog(ctx *gin.Context) {

	body := blog.NewBody()

	// 这和业务无关的错误
	if err := ctx.BindJSON(body); err != nil {
		// 处理异常
		ctx.JSON(http.StatusBadRequest, utils.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}

	// 和业务有关了
	if err := body.Validate(); err != nil {
		// 处理异常
		if e, ok := err.(*utils.APIError); ok {
			ctx.JSON(e.HttpStatus, e)
		} else {
			ctx.JSON(http.StatusInternalServerError, utils.NewAPIError(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	// 调用业务的 CreateBlog
	ins, err := h.svc.CreateBlog(ctx.Request.Context(), body)
	if err != nil {
		// 处理异常
		if e, ok := err.(*utils.APIError); ok {
			ctx.JSON(e.HttpStatus, e)
		} else {
			ctx.JSON(http.StatusInternalServerError, utils.NewAPIError(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	ctx.JSON(http.StatusOK, ins)
}

func (h *handler) QueryBlog(ctx *gin.Context) {
	in := blog.NewQueryBlogRequest()

	// GET 请求，从 url query 中获取参数
	pageSize, err := strconv.ParseInt(ctx.DefaultQuery("pageSize", "20"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewAPIError(http.StatusBadRequest, err.Error()))
	}
	in.PageSize = int(pageSize)

	pageNumber, err := strconv.ParseInt(ctx.DefaultQuery("pageNumber", "1"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewAPIError(http.StatusBadRequest, err.Error()))
	}
	in.PageNumber = int(pageNumber)

	in.Author = ctx.Query("author")
	in.Keywords = ctx.Query("keyword")

	set, err := h.svc.QueryBlog(ctx, in)
	if err != nil {
		// 处理异常
		if e, ok := err.(*utils.APIError); ok {
			ctx.JSON(e.HttpStatus, e)
		} else {
			ctx.JSON(http.StatusInternalServerError, utils.NewAPIError(http.StatusInternalServerError, err.Error()))
		}
		return
	}
	ctx.JSON(http.StatusOK, set)
}

func (h *handler) DescribeBlog(ctx *gin.Context) {
	in := &blog.DescribeBlogRequest{}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}
	in.Id = int(id)

	ins, err := h.svc.DescribeBlog(ctx, in)
	if err != nil {
		// 处理异常
		if e, ok := err.(*utils.APIError); ok {
			ctx.JSON(e.HttpStatus, e)
		} else {
			ctx.JSON(http.StatusInternalServerError, utils.NewAPIError(http.StatusInternalServerError, err.Error()))
		}
		return
	}
	ctx.JSON(http.StatusOK, ins)
}

func (h *handler) UpdateBlog(ctx *gin.Context) {
	in := blog.NewUpdateBlogRequest()

	if err := ctx.BindJSON(in); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}

	ins, err := h.svc.UpdateBlog(ctx, in)
	if err != nil {
		// 处理异常
		if e, ok := err.(*utils.APIError); ok {
			ctx.JSON(e.HttpStatus, e)
		} else {
			ctx.JSON(http.StatusInternalServerError, utils.NewAPIError(http.StatusInternalServerError, err.Error()))
		}
		return
	}
	ctx.JSON(http.StatusOK, ins)
}

func (h *handler) DeleteBlog(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))
	in := blog.NewDeleteBlogRequest(id)

	ins, err := h.svc.DeleteBlog(ctx, in)
	if err != nil {
		// 处理异常
		if e, ok := err.(*utils.APIError); ok {
			ctx.JSON(e.HttpStatus, e)
		} else {
			ctx.JSON(http.StatusInternalServerError, utils.NewAPIError(http.StatusInternalServerError, err.Error()))
		}
		return
	}
	ctx.JSON(http.StatusOK, ins)
}
