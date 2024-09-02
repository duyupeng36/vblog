package impl_test

import (
	"vblog/apps/blog"
	"vblog/apps/blog/impl"
	"vblog/conf"
)

var (
	controller blog.Service
)

func init() {
	var err error

	err = conf.LoadCOnfigFromEnv()
	if err != nil {
		panic(err)
	}

	controller, err = impl.NewBlogService()
	if err != nil {
		panic(err)
	}
}
