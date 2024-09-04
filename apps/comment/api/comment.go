package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vblog/apps/comment"
	"vblog/utils"
)

func (h *Handler) CreateComment(ctx *gin.Context) {
	body := comment.NewBody()

	// 获取用户传递的参数
	if err := ctx.BindJSON(body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewAPIError(http.StatusBadRequest, err.Error()))
	}

	// 调用业务接口
	ins, err := h.svc.CreateComment(ctx, body)
	if err != nil {
		if e, ok := err.(*utils.APIError); ok {
			ctx.JSON(e.HttpStatus, e)
		} else {
			ctx.JSON(http.StatusInternalServerError, utils.NewAPIError(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	ctx.JSON(http.StatusOK, ins)
}
