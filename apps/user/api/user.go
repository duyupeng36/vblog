package api

import (
	"net/http"
	"vblog/apps/user"
	"vblog/utils"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateUser(ctx *gin.Context) {
	in := user.NewCreateUserRequest()

	if err := ctx.BindJSON(in); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}

	ins, err := h.svc.CreateUser(ctx.Request.Context(), in)
	if err != nil {
		if e, ok := err.(*utils.APIError); ok {
			ctx.JSON(e.HttpStatus, e)
		} else {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, ins)
}
