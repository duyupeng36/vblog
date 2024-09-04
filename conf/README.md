# 项目配置

项目依赖的所有参数都将从配置对象 `Config` 中读取

本项目通过两种方式加载配置
+ 通过文件的方式提供配置：采用 [toml](https://toml.io/cn) 格式的配置文件
+ 通过环境变量的方式加载配置，采用 [env](https://github.com/caarlos0/env) 包加载
+ [加载配置的实现](./load.go)


