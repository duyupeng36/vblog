package api

import (
	"errors"
	"vblog/apps/comment"
	"vblog/ioc"
	"vblog/protocal/middleware"

	"github.com/gin-gonic/gin"
)

// Handler 负责实现具体的 API
type Handler struct {
	// 需要一个业务的具体实现
	svc comment.Service
}

func (h *Handler) Init() error {
	svc, ok := ioc.Container().GetController(comment.AppName).(comment.Service)

	if !ok {
		return errors.New("get controller faild")
	}
	h.svc = svc
	return nil
}

func (h *Handler) Name() string {
	return comment.AppName
}

// 还需要将 API 与 HTTP 路由对应上

func (h *Handler) Registry(r gin.IRouter) {
	r.Use(middleware.Auth)
	// 注册 路由
	r.POST("", h.CreateComment)
}

func init() {
	ioc.Container().RegistryHandler(&Handler{})
}
