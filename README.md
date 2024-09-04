# 微博客

## RESTful API

REST 与技术无关，代表的是一种软件架构风格，REST 是 **Representational State Transfer** 的简称，中文翻译为 “**表征状态转移**” 或 “表现层状态转化”

> 状态：_资源_ 的状态
>
> 资源就是网络上的实体。可以是文本、图片、音频、视频等
>
> **资源总是以一定的格式来表现自己**

**URI** 作为资源标识符，**HTTP 请求方法** 表征对这些资源的操作

> 简单讲
>
> URL 表示资源，HTTP 请求方法表示资源状态的改变
>

RESTful API 设计中使用 $4$ 个 HTTP 请求方法资源 **状态的变化**

|   请求方法   | 动作     |
|:--------:|:-------|
|  `GET`   | 获取资源   |
|  `POST`  | 新建资源   |
|  `PUT`   | 更新资源   |
| `PATCH`  | 部分更新资源 |
| `DELETE` | 删除资源   |

> [!tip] 状态的变化
>
> **幂等性**：无论一个操作被执行一次还是多次，执行后的效果都是相同的，例如 `GET`
>
> `POST` `DELETE` `PUT` `PATCH` 都对应资源的各自状态变化


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
├── README.md.md
├── apps                          # 存放所有业务的目录
│   ├── README.md.md
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
│   ├── README.md.md
│   ├── root.go
│   └── start
│       └── start.go
├── conf                        # 项目的配置文件
│   ├── README.md.md
│   ├── config.go
│   ├── config_test.go
│   ├── http.go
│   ├── load.go
│   ├── load_test.go
│   └── mysql.go
├── docs
├── etc                        # 配置文件
│   ├── README.md.md
│   ├── config.env
│   ├── config.example.env
│   ├── config.toml
│   └── unit_test.env
├── go.mod
├── go.sum
├── main.go                   # 程序的入口文件
├── protocal                  # 协议服务器，对外提供服务
│   ├── README.md.md
│   └── http.go
├── utils                     # 全局工具
│   ├── README.md.md
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

## 控制反转(Inversion of Control, ioc)

目前，项目的业务逻辑需要在项目启动的时候手动添加业务逻辑的初始化。在 [start](./cmd/start/start.go)，需要写一长串的初始化代码，项目接口才能被使用

每个业务都要在这里初始化。并且每个业务的 `Handler` 都直接依赖 `Service` 的实现，导致
+ 注意里面不能集中在写业务功能上,  **不希望在 `apps` 模块之外还要做业务代码的添加**
+ **对象的依赖问题**, 一些其他人的负责的业务, 可能会已很多业务对象来开发上层业务

```go
// 初始化业务控制器
blogServer, err := impl.NewBlogService()
if err != nil {
    return err
}
commentServer, err := impl2.NewCommentService()
if err != nil {
    return err
}
// 初始化 Http 控制器
blogHttpHandler := api.NewHandler(blogServer)
commentHttpHandler := api2.NewHandler(commentServer)

// 创建 Gin 路由引擎
routerEngine := gin.Default()

// 注册路由
blogHttpHandler.Registry(routerEngine)
commentHttpHandler.Registry(routerEngine)

// 创建 Http 服务端
httpServer := protocal.NewHttp(routerEngine)
```

[业务间对象依赖问题](./docs/object.depence.drawio)，当对象间 **直接依赖** 的数量变得越来越多，从而导致依赖关系非常难以维护。

为了不让对象直接依赖，我们采用 **控制反转(Inversion of Control, ioc)**：是一种是面向对象编程中的一种设计原则，用来减低计算机代码之间的耦合度。其基本思想是：**借助于“第三方”实现具有依赖关系的对象之间的解耦**

[ioc](./docs/ioc.drawio)

采用 ioc 之后，业务代码只需要
1. 先写业务逻辑
2. 再写接口 RESTful API
3. 实例注册到 ioc 中

```go
package apps

import (
	// 注册 handler
	_ "vblog/apps/blog/api"

	// 注册业务具体实现
	_ "vblog/apps/blog/impl"
)
```

4. 启动时的时候，注册所有实例：`apps/registry.go`

```go 
_ "vblog/apps" // 导入 apps
```

## 启动

```shell
./vblog start
```
