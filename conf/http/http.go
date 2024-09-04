package http

import "fmt"

type Http struct {
	Host         string `json:"host" toml:"host" env:"HTTP_HOST"`
	Port         int    `json:"port" toml:"port" env:"HTTP_PORT"`
	ReadTimeout  int    `json:"readTimeout" toml:"readTimeout"  env:"HTTP_READTIMEOUT"`
	WriteTimeout int    `json:"writeTimeout" toml:"writeTimeout" env:"HTTP_WRITETIMEOUT"`
}

// Default 提供 HTTP 服务的默认配置
func Default() *Http {
	return &Http{
		Host:         "localhost",
		Port:         8080,
		ReadTimeout:  5,
		WriteTimeout: 5,
	}
}

// Addr 获取服务监听的地址和端口
func (h *Http) Addr() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}
