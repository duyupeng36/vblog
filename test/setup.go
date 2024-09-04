package test

import (
	"vblog/conf"
	"vblog/ioc"
)

// DevelopmentSetUp，初始化 IOC
func DevelopmentSetUp() {
	var err error

	err = conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	// 初始化 ioc

	ioc.Container().Init(nil)

}
