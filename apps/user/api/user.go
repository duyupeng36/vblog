package api

import (
	"net/http"
	"strconv"
	"strings"
	"vblog/apps/user"
	"vblog/utils/errors"
	"vblog/utils/response"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateUser(ctx *gin.Context) {
	in := user.NewCreateUserRequest()

	if err := ctx.BindJSON(in); err != nil {
		response.SendFailed(ctx, errors.NewBadRequestError(err.Error()))
		return
	}

	ins, err := h.svc.CreateUser(ctx.Request.Context(), in)
	if err != nil {

		return
	}
	response.SendSuccess(ctx, ins)
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
		response.SendFailed(ctx, err)
		return
	}
	response.SendSuccess(ctx, ins)
}

func (h *handler) Login(ctx *gin.Context) {
	in := user.NewLoginRequest()

	if err := ctx.BindJSON(in); err != nil {
		response.SendFailed(ctx, errors.NewBadRequestError(err.Error()))
		return
	}

	tk, err := h.svc.Login(ctx, in)

	if err != nil {
		response.SendFailed(ctx, err)
		return
	}

	ctx.SetCookie(user.AUTOH_COOKIE_NAME, tk.AccessToken, 3600*12, "/", strings.TrimRight(ctx.Request.Host, ":0123456789"), false, false)

	response.SendSuccess(ctx, tk)
}
