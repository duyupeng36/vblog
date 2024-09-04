package ioc

import "github.com/gin-gonic/gin"

var (
	container *ioc
)

func Container() *ioc {
	return container
}

type IoController interface {
	Init() error  // 用于初始化对象
	Name() string // 获取对象的名字
}

type IoCHandler interface {
	IoController
	Registry(r gin.IRouter) // 注册路由
}

type ioc struct {
	controller map[string]IoController // 管理 Controller 对象
	handler    map[string]IoCHandler   // 管理 Handler 对象
}

func New() *ioc {
	return &ioc{
		controller: make(map[string]IoController, 20),
		handler:    make(map[string]IoCHandler, 20),
	}
}

// Init 方法，用于初始化 IOC 管理的对象
func (i *ioc) Init(r gin.IRouter) error {
	for _, v := range i.controller {
		if err := v.Init(); err != nil {
			return err
		}
	}

	for _, v := range i.handler {
		if err := v.Init(); err != nil {
			return err
		}
		v.Registry(r)
	}
	return nil
}

func init() {
	container = New()
}
