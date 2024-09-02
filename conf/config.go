package conf

import (
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	MySQL *MySQL `toml:"mysql" json:"mysql"`
	Http  *Http  `json:"http" toml:"http"`
}

// DefaultConfig 返回一个据哟默认配置的 Config 对象
func DefaultConfig() *Config {
	return &Config{
		MySQL: defaultMySQL(),
		Http:  defaultHttp(),
	}
}

func (c *Config) String() string {
	result, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return ""
	}
	return string(result)
}
