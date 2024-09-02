package conf

import (
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/pelletier/go-toml/v2"
)

var (
	config *Config
)

// 获取全局配置
func C() *Config {

	if config == nil {
		log.Fatalf("the config doesn't load, please call the func LoadConfigFromToml or LoadCOnfigFromEnv")
	}
	return config
}

// LoadConfigFromToml 从配置文件中读取
func LoadConfigFromToml(path string) error {

	// 首先，创建一个默认配置对象
	defaultConfig := DefaultConfig()

	// 然后，从配置文件中读取，并覆盖默认配置
	f, err := os.Open(path)
	if err != nil {
		// 配置文件不存在，返回默认配置
		if os.IsNotExist(err) {
			log.Println("the config.toml is not exist, will be use the default config")
			config = defaultConfig
			return nil
		}
		log.Printf("open config.toml failed: %v\n", err)
		return err
	}

	defer f.Close()

	// 使用 toml 配置文件覆盖默认配置
	err = toml.NewDecoder(f).Decode(defaultConfig)
	if err != nil {
		log.Printf("Decode the config.toml failed: %v\n", err)
		return err
	}

	config = defaultConfig

	return nil
}

// LoadCOnfigFromEnv 加载环境变量的配置
// 使用 "github.com/caarlos0/env/v6" 包
func LoadConfigFromEnv() error {
	defaultconfig := DefaultConfig()

	if err := env.Parse(defaultconfig); err != nil {
		return err
	}

	config = defaultconfig
	return nil
}
