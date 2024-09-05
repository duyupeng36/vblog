package impl_test

import (
	"context"
	"vblog/apps/user"
	"vblog/ioc"
	"vblog/test"
)

var (
	controller user.Service
	ctx        = context.Background()
)

func init() {
	test.DevelopmentSetUp()
	controller = ioc.Container().GetController(user.AppName).(user.Service)
}
