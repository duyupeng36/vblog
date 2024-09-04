package impl_test

import (
	"vblog/apps/blog"
	"vblog/ioc"
	"vblog/test"
)

var (
	controller blog.Service
)

func init() {
	test.DevelopmentSetUp()                                                 // 初始化测试环境
	controller = ioc.Container().GetController(blog.AppName).(blog.Service) // 从 ioc 获取 controller
}

// func init() {
// 	// var err error

// 	// err = conf.LoadConfigFromEnv()
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// controller, err = impl.NewBlogService()
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// }
