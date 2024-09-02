package blog

import (
	"context"
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
type QueryBlogRequest struct {
}

// UpdateBlogRequest 更新博客的请求结构
type UpdateBlogRequest struct {
}

// DescribeBlogRequest 查询一篇博客的请求结构体
type DescribeBlogRequest struct {
}

// DeleteBlogRequest 删除一篇博客的请求结构体
type DeleteBlogRequest struct {
}
