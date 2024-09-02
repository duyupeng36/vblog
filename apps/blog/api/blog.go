package api

import (
	"net/http"
	"vblog/apps/blog"
	"vblog/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateBlog(ctx *gin.Context) {

	body := blog.NewBody()

	// 这和业务无关的错误
	if err := ctx.BindJSON(body); err != nil {
		// 处理异常
		ctx.JSON(http.StatusBadRequest, utils.NewAPIError(http.StatusBadRequest, err.Error()))

	}

	// 和业务有关了
	if err := body.Validate(); err != nil {
		// 处理异常
		if e, ok := err.(*utils.APIError); ok {
			ctx.JSON(e.HttpStatus, e)
		} else {
			ctx.JSON(http.StatusInternalServerError, utils.NewAPIError(http.StatusInternalServerError, err.Error()))
		}
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
	}

	ctx.JSON(http.StatusOK, ins)
}
