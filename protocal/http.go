package protocal

import (
	"context"
	"net/http"
	"time"
	"vblog/conf"
	"vblog/logger"

	"github.com/gin-gonic/gin"
)

type Http struct {
	server *http.Server
}

func NewHttp(r *gin.Engine) *Http {
	return &Http{
		server: &http.Server{
			Handler: r, // Gin 路由引擎

			// Http 服务配置，先入不能这样硬编码，而是应该从配置对象上读取
			Addr: conf.C().Http.Addr(), // "127.0.0.1:8080",
			// 超时时间
			ReadTimeout:  time.Duration(conf.C().Http.ReadTimeout) * time.Second,  // 5 * time.Second
			WriteTimeout: time.Duration(conf.C().Http.WriteTimeout) * time.Second, // 5 * time.Second
		},
	}
}

// Start 启动服务
func (h *Http) Start() error {
	// 启动http服务
	logger.Logger().Info().Msgf("start http server at %s", h.server.Addr)
	return h.server.ListenAndServe()
}

func (h *Http) Stop(ctx context.Context) error {
	// Shatdown: 首先关闭监听，然后等待已有连接处理完成
	return h.server.Shutdown(ctx)
}
