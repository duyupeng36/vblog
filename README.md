# 微博客

```shell
vblog   # 项目目录
├── apps # 存放所有业务的目录
│   ├── blog  # 博客管理业务
│   │   ├── api  # 实现各种协议服务的 API
│   │   │   ├── blog.go
│   │   │   └── http.go
│   │   ├── const.go
│   │   ├── impl  # 实现业务功能
│   │   │   ├── blog.go
│   │   │   ├── blog_test.go
│   │   │   ├── dao.go
│   │   │   ├── impl.go
│   │   │   └── impl_test.go
│   │   ├── interface.go  # 业务接口定义
│   │   └── model.go  # 业务模型
│   ├── comment  # 博客评论管理
│   │   ├── api
│   │   │   ├── comment.go
│   │   │   └── http.go
│   │   ├── impl
│   │   │   ├── comment.go
│   │   │   ├── comment_test.go
│   │   │   ├── impl.go
│   │   │   └── impl_test.go
│   │   ├── interface.go
│   │   └── model.go
│   └── user  # 用户管理
│       └── model.go
├── cmd  # 项目的 CLI
│   └── README.md
├── conf # 项目的配置管理
│   ├── README.md
│   ├── config.go
│   └── config_test.go
├── docs
├── etc  # 项目的配置文件
│   ├── README.md
│   ├── config.env
│   ├── config.example.env
│   ├── config.toml
│   └── unit_test.env
├── go.mod
├── go.sum
├── main.go # 项目启动的入口文件
└── protocal  # 项目对外提供的协议服务
    └── README.md
```

