package api

import (
	"vblog/apps/blog"

	"github.com/gin-gonic/gin"
)

// NewHandler 初始化一些基础数据
// 具体传递那个实现，在程序启动初始化时
func NewHandler(svc blog.Service) *Handler {

	return &Handler{
		svc: svc,
	}
}

// Handler 负责实现具体的 API
type Handler struct {
	// 需要一个业务的具体实现
	svc blog.Service
}

// 还需要将 API 与 HTTP 路由对应上

func (h *Handler) Registry(r gin.IRouter) {

	// 注册 路由
	r.POST("/vblog/api/v1/blogs", h.CreateBlog)
}
