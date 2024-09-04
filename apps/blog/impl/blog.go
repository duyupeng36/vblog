package impl

import (
	"context"
	"fmt"
	"net/http"
	"vblog/apps/blog"
	"vblog/utils"
)

const (
	StatusOK          = 200 // 正常
	StatusParamsError = 300 // 请求参数错误
	StatusStoreError  = 400 // 入库失败
)

// CreateBlog 创建博客
func (i *impl) CreateBlog(ctx context.Context, body *blog.Body) (*blog.Blog, error) {
	// 验证请求参数
	if err := body.Validate(); err != nil {
		return nil, utils.NewAPIError(StatusParamsError, err.Error()).SetHttpStatus(http.StatusBadRequest)
	}

	// 验证成功
	// 创建一个 blog 对象
	// 直接使用结构体创建对象：这是不合适的，因为后续可能
	// ins := blog.Blog{
	// 	Meta: &blog.Meta{},
	// 	Body: body,
	// }
	// ins.Body.Extra 当我们添加了额外字段，如果是一个 map ，没有初始化可能会造成程序错误
	// 使得代码兼容性较差

	ins := blog.NewBlog(body)

	// 调用 orm 将其持久化到数据库中
	// 注意，用户可能随时取消请求，当用户取消请求后，数据库操作也应该取消
	// ctx 就是通知 orm 取消这个数据库操作的
	// i.db.Create(ins)
	if err := i.db.WithContext(ctx).Create(ins).Error; err != nil {
		return nil, utils.NewAPIError(StatusStoreError, err.Error()).SetHttpStatus(http.StatusInternalServerError)
	}

	return ins, nil
}

// QueryBlog 获取文章列表
func (i *impl) QueryBlog(ctx context.Context, in *blog.QueryBlogRequest) (*blog.BlogSet, error) {

	var set = blog.NewBlogSet()

	// 构建 SQL 语句
	query := i.db.WithContext(ctx).Model(&blog.Blog{})

	if in.Keywords != "" {
		query = query.Where("content LIKE ?", "%"+in.Keywords+"%")
	}

	if in.Author != "" {
		query = query.Where("author = ?", in.Author)
	}

	// 分页
	if err := query.Count(&set.Total).Offset(in.Offset()).Limit(in.PageSize).Find(&set.Items).Error; err != nil {
		return nil, utils.NewAPIError(http.StatusBadRequest, err.Error())
	}

	return set, nil
}

// DescribeBlog 获取一篇博客
func (i *impl) DescribeBlog(ctx context.Context, in *blog.DescribeBlogRequest) (*blog.Blog, error) {

	result := &blog.Blog{}

	// 构建 SQL 语句
	query := i.db.WithContext(ctx).Model(&blog.Blog{})

	if in.Id == 0 {
		return nil, utils.NewAPIError(http.StatusBadRequest, fmt.Errorf("query Id must give").Error())
	}

	if err := query.Take(&result, in.Id).Error; err != nil {
		return nil, utils.NewAPIError(http.StatusNotFound, err.Error())
	}

	return result, nil
}

// UpdateBlog 修改一篇博客
func (i *impl) UpdateBlog(context.Context, *blog.UpdateBlogRequest) (*blog.Blog, error) {
	return nil, nil
}

// DeleteBlog 删除博客，返回删除的对象，用于前端展示或对象最终
func (i *impl) DeleteBlog(context.Context, *blog.DeleteBlogRequest) (*blog.Blog, error) {
	return nil, nil
}
