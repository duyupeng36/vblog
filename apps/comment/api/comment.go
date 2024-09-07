package api

import (
	"vblog/apps/comment"
	"vblog/utils/errors"
	"vblog/utils/response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComment(ctx *gin.Context) {
	body := comment.NewBody()

	// 获取用户传递的参数
	if err := ctx.BindJSON(body); err != nil {
		response.SendFailed(ctx, errors.NewBadRequestError(err.Error()))
		return
	}

	// 调用业务接口
	ins, err := h.svc.CreateComment(ctx, body)
	if err != nil {
		response.SendFailed(ctx, err)
		return
	}
	response.SendSuccess(ctx, ins)
}
