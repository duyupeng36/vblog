package start

import (
	"vblog/apps/blog/api"
	"vblog/apps/blog/impl"
	"vblog/conf"
	"vblog/protocal"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var configType string
var configPath string

var Cmd = &cobra.Command{
	Use:   "start",
	Short: "start vblog http server",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 组装服务

		// 读取配置文件
		switch configType {
		case "env":
			if err := conf.LoadCOnfigFromEnv(); err != nil {
				return err
			}
		default:
			if err := conf.LoadConfigFromToml(configPath); err != nil {
				return err
			}
		}

		// 初始化业务控制器
		blogServer, err := impl.NewBlogService()
		if err != nil {
			return err
		}
		// 初始化 Http 控制器
		blogHtppHandler := api.NewHandler(blogServer)
		// 创建 Gin 路由引擎
		routerEngine := gin.Default()
		// 注册路由
		blogHtppHandler.Registry(routerEngine)

		// 创建 Http 服务端
		httpServer := protocal.NewHttp(routerEngine)
		// 启动服务
		if err := httpServer.Start(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	Cmd.Flags().StringVarP(&configType, "config-type", "t", "file", "How to read config")
	Cmd.Flags().StringVarP(&configPath, "config-file", "f", "etc/config.toml", "Where read config")
}
