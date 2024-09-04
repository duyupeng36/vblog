package conf

import (
	"io"
	"os"
	"vblog/logger"

	"github.com/caarlos0/env/v6"
	"github.com/pelletier/go-toml/v2"
)

// LoadConfigFromToml 从配置文件中读取
func LoadConfigFromToml(path string) error {

	logger.Logger().Info().Msgf("Loading config from %s", path)
	// 首先，创建一个默认配置对象
	defaultConfig := DefaultConfig()

	// 然后，从配置文件中读取，并覆盖默认配置
	f, err := os.Open(path)
	if err != nil {
		// 配置文件不存在，返回默认配置
		if os.IsNotExist(err) {
			logger.Logger().Info().Msgf("the %s is not exist, will be use the default config", path)
			config = defaultConfig
			return nil
		}
		logger.Logger().Fatal().Msgf("open %s failed: %v\n", path, err)
		return err
	}

	defer func(f io.Closer) {
		err := f.Close()
		if err != nil {
			logger.Logger().Warn().Msgf("close %s failed: %v\n", path, err)
		}
	}(f)

	// 使用 toml 配置文件覆盖默认配置
	err = toml.NewDecoder(f).Decode(defaultConfig)
	if err != nil {
		logger.Logger().Fatal().Msgf("Decode the %s failed: %v\n", path, err)
		return err
	}

	config = defaultConfig

	logger.Logger().Info().Msg("Finished loading config")
	return nil
}

// LoadConfigFromEnv 加载环境变量的配置
// 使用 "github.com/caarlos0/env/v6" 包
func LoadConfigFromEnv() error {
	logger.Logger().Info().Msg("Loading config from environment")
	defaultConfig := DefaultConfig()

	if err := env.Parse(defaultConfig); err != nil {
		logger.Logger().Fatal().Msgf("load default config failed: %v\n", err)
		return err
	}

	config = defaultConfig
	logger.Logger().Info().Msg("Finished loading config")
	return nil
}
