package api

import (
	"errors"
	"vblog/apps/blog"
	"vblog/ioc"

	"github.com/gin-gonic/gin"
)

// NewHandler 初始化一些基础数据
// 具体传递那个实现，在程序启动初始化时
// 不用向这样直接依赖，而是通过 ioc 获取
// func NewHandler(svc blog.Service) *handler {

//		return &handler{
//			svc: svc,
//		}
//	}

// handler 负责实现具体的 API
type handler struct {
	// 需要一个业务的具体实现
	svc blog.Service
}

func (h *handler) Init() error {

	svc, ok := ioc.Container().GetController(blog.AppName).(blog.Service)

	if !ok {
		return errors.New("get controller faild")
	}
	h.svc = svc
	return nil
}

// Registry 将 API 与 HTTP 路由对应上
func (h *handler) Registry(r gin.IRouter) {

	// 注册 路由
	r.POST("/vblog/api/v1/blogs", h.CreateBlog)
	r.GET("/vblog/api/v1/blogs", h.QueryBlog)
	r.GET("/vblog/api/v1/blogs/:id", h.DescribeBlog)
	r.PATCH("/vblog/api/v1/blogs", h.UpdateBlog)
	r.DELETE("/vblog/api/v1/blogs/:id", h.DeleteBlog)
}

func (h *handler) Name() string {
	return blog.AppName
}

// 通过 import
func init() {
	ioc.Container().RegistryHandler(&handler{})
}
