package impl_test

import (
	"vblog/apps/comment"
	"vblog/apps/comment/impl"
	"vblog/conf"
)

var (
	controller comment.Service
)

func init() {
	var err error

	err = conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	controller, err = impl.NewCommentService()
	if err != nil {
		panic(err)
	}
}
