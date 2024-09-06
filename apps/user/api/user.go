package api

import (
	"net/http"
	"strconv"
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

func (h *handler) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	in := user.NewDeleteUserRequest(int(id))

	ins, err := h.svc.DeleteUser(ctx.Request.Context(), in)
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

func (h *handler) Login(ctx *gin.Context) {
	in := user.NewLoginRequest()

	if err := ctx.BindJSON(in); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}

	tk, err := h.svc.Login(ctx, in)

	if err != nil {
		if e, ok := err.(*utils.APIError); ok {
			ctx.JSON(e.HttpStatus, e)
		} else {
			ctx.JSON(http.StatusInternalServerError, utils.NewAPIError(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	ctx.SetCookie(user.AUTOH_COOKIE_NAME, tk.AccessToken, 3600*12, "/", ctx.Request.Host, false, false)

	ctx.JSON(http.StatusOK, tk)
}
