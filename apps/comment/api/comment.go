package api

import (
	"github.com/gin-gonic/gin"
	"vblog/apps/comment"
)

func (h *Handler) CreateComment(ctx *gin.Context) {
	ctx.BindJSON(&comment.Comment{})
}
