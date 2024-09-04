# 微博客

## 项目代码的组织

+ [配置管理(conf包)](./conf/README.md) 提供项目运行的参数。例如，数据库的连接，项目运行的端口等
+ [业务代码(apps包)](./apps/README.md) 业务代码按照 DDD 模式编写
+ [项目文档(docs目录)](./docs/README.md) 项目开发过程中的文档等
+ [项目提供的服务(protocol包)](./protocal/README.md) 项目提供的 HTTP，RPC 接口的服务
+ [项目的CLI(cmd包)](./cmd/README.md) 提供项目命令。例如，`vblog strat/init` 等命令
+ [项目的入口(main.go)](./main.go) 启动项目的入口
+ [项目的配置文件(etc目录)](./etc/README.md) 项目的配置文件

```shell
vblog   # 项目目录
├── README.md
├── apps                          # 存放所有业务的目录
│   ├── README.md
│   ├── blog                      # 博客管理业务
│   │   ├── api                     # 实现各种协议服务的 API
│   │   │   ├── blog.go
│   │   │   └── http.go
│   │   ├── const.go
│   │   ├── impl                    # 实现业务功能
│   │   │   ├── blog.go
│   │   │   ├── blog_test.go
│   │   │   ├── dao.go
│   │   │   ├── impl.go
│   │   │   └── impl_test.go
│   │   ├── interface.go            # 业务接口定义
│   │   └── model.go                # 业务模型
│   ├── comment                   # 博客评论管理
│   │   ├── api
│   │   │   ├── comment.go
│   │   │   └── http.go
│   │   ├── impl
│   │   │   ├── comment.go
│   │   │   ├── comment_test.go
│   │   │   ├── impl.go
│   │   │   └── impl_test.go
│   │   ├── interface.go
│   │   └── model.go
│   └── user                     # 用户管理
│       └── model.go
├── cmd                          # 项目的 CLI
│   ├── README.md
│   ├── root.go
│   └── start
│       └── start.go
├── conf                        # 项目的配置文件
│   ├── README.md
│   ├── config.go
│   ├── config_test.go
│   ├── http.go
│   ├── load.go
│   ├── load_test.go
│   └── mysql.go
├── docs
├── etc                        # 配置文件
│   ├── README.md
│   ├── config.env
│   ├── config.example.env
│   ├── config.toml
│   └── unit_test.env
├── go.mod
├── go.sum
├── main.go                   # 程序的入口文件
├── protocal                  # 协议服务器，对外提供服务
│   ├── README.md
│   └── http.go
├── utils                     # 全局工具
│   ├── README.md
│   ├── error.go
│   └── validate.go
```

## 开发流程

1. 完成项目配置加载并测试 
2. 开始编写业务代码，测试 
3. 业务逻辑测试通过后，开始编写业务对外提供的 **HTTP 接口** 或 **RPC 接口**
4. 编写项目的 **接口服务**，提供启动服务的方法
5. 编写项目的 CLI 命令
6. 启动项目

## 启动

```shell
./vblog start
```

