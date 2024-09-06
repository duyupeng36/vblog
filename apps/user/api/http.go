package api

import (
	"errors"
	"vblog/apps/user"
	"vblog/ioc"
	"vblog/protocal/middleware"

	"github.com/gin-gonic/gin"
)

type handler struct {
	svc user.Service
}

func (h *handler) Init() error {
	svc, ok := ioc.Container().GetController(user.AppName).(user.Service)
	if !ok {
		return errors.New("get controller faild")
	}
	h.svc = svc
	return nil
}

func (h *handler) Name() string {
	return user.AppName
}

func (h *handler) Registry(r gin.IRouter) {
	r.POST("", h.CreateUser)
	r.POST("/token", h.Login)
	auth := r.Group("").Use(middleware.Auth)
	{
		auth.DELETE("/:id", h.DeleteUser)
	}
}

func init() {
	ioc.Container().RegistryHandler(&handler{})
}
