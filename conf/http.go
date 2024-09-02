package conf

import "fmt"

type Http struct {
	Host         string `json:"host" toml:"host" env:"HTTP_HOST"`
	Port         int    `json:"port" toml:"port" env:"HTTP_PORT"`
	ReadTimeout  int    `json:"readTimeout" toml:"readTimeout"  env:"HTTP_READTIMEOUT"`
	WriteTimeout int    `json:"writeTimeout" toml:"writeTimeout" env:"HTTP_WRITETIMEOUT"`
}

func defaultHttp() *Http {
	return &Http{
		Host:         "localhost",
		Port:         8080,
		ReadTimeout:  5,
		WriteTimeout: 5,
	}
}

func (h *Http) Addr() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}
