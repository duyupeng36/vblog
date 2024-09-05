package blog

import (
	"context"
)

const (
	// 业务的名字
	AppName = "blogs"
)

// 博客管理的业务接口声明。隔离服务的具体实现

type Service interface {
	// QueryBlog 获取文章列表
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	// CreateBlog 创建博客
	CreateBlog(context.Context, *Body) (*Blog, error)
	// DescribeBlog 获取一篇博客
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	// UpdateBlog 修改一篇博客
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)
	// DeleteBlog 删除博客，返回删除的对象，用于前端展示或对象最终
	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)
}

// QueryBlogRequest 查询博客列表的请求结构体
// 至少要：服务端分页，默认 1页20个
// 关键字查询，模糊查询
// 条件过滤，作者是谁
type QueryBlogRequest struct {
	PageSize   int    // 每页的大小
	PageNumber int    // 当前页
	Keywords   string // 模糊查找，通过文章内容
	Author     string // 条件，作者是谁
}

func (q *QueryBlogRequest) Offset() int {
	return (q.PageNumber - 1) * q.PageSize
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

// UpdateBlogRequest 更新博客的请求结构
type UpdateBlogRequest struct {
	Id      int               `json:"id"`
	Title   string            `json:"title"`
	Content string            `json:"content"`
	Tags    map[string]string `json:"tags"`
}

func NewUpdateBlogRequest() *UpdateBlogRequest {
	return &UpdateBlogRequest{}
}

// DescribeBlogRequest 查询一篇博客的请求结构体
type DescribeBlogRequest struct {
	Id int // blog id
}

func NewDescribeBlogRequest(id int) *DescribeBlogRequest {
	return &DescribeBlogRequest{
		Id: id,
	}
}

// DeleteBlogRequest 删除一篇博客的请求结构体
type DeleteBlogRequest struct {
	Id int
}

func NewDeleteBlogRequest(id int) *DeleteBlogRequest {
	return &DeleteBlogRequest{
		Id: id,
	}
}
