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

func TestDescribeBlog(t *testing.T) {
	req := blog.NewDescribeBlogRequest(1)
	ins, err := controller.DescribeBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDeleteBlog(t *testing.T) {
	req := blog.NewDeleteBlogRequest(1)
	ins, err := controller.DeleteBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestUpdateBlog(t *testing.T) {
	req := blog.NewUpdateBlogRequest()
	req.Id = 2
	req.Content = "Python"
	req.Tags = map[string]string{
		"ca": "ci",
	}
	req.Title = "描述器"

	ins, err := controller.UpdateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
