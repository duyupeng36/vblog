package impl_test

import (
	"context"
	"testing"
	"vblog/apps/blog"
)

var (
	ctx = context.Background()
)

func TestCreateBlog(t *testing.T) {
	in := blog.NewBody()
	in.Title = "vblog_test"
	in.Author = "ddd"
	in.Content = "test_test"
	ins, err := controller.CreateBlog(ctx, in)
	if err != nil {
		t.Error(err)
	}

	t.Log(ins)
}

func TestQueryBlog(t *testing.T) {
	req := blog.NewQueryBlogRequest()
	req.PageSize = 1
	req.Keywords = "项目"
	set, err := controller.QueryBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)

}
