package conf

import (
	"encoding/json"
	"log"
	"vblog/conf/http"
	"vblog/conf/mysql"
)

var (
	config *Config // 全局配置对象
)

// C 获取全局配置
func C() *Config {

	if config == nil {
		log.Fatalf("the config doesn't load, please call the func LoadConfigFromToml or LoadCOnfigFromEnv")
	}
	return config
}

type Config struct {
	MySQL *mysql.MySQL `toml:"mysql" json:"mysql"`
	Http  *http.Http   `json:"http" toml:"http"`
}

// DefaultConfig 返回一个据哟默认配置的 Config 对象
func DefaultConfig() *Config {
	return &Config{
		MySQL: mysql.Default(),
		Http:  http.Default(),
	}
}

func (c *Config) String() string {
	result, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return ""
	}
	return string(result)
}
