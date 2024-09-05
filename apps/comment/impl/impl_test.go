package impl_test

import (
	"vblog/apps/comment"
	"vblog/ioc"
	"vblog/test"
)

var (
	controller comment.Service
)

func init() {
	test.DevelopmentSetUp()                                                       // 初始化测试环境
	controller = ioc.Container().GetController(comment.AppName).(comment.Service) // 从 ioc 获取 controller
}

// func init() {
// 	var err error

// 	err = conf.LoadConfigFromEnv()
// 	if err != nil {
// 		panic(err)
// 	}

// 	controller, err = impl.NewCommentService()
// 	if err != nil {
// 		panic(err)
// 	}
// }
