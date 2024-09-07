package api

import (
	"strconv"
	"vblog/apps/blog"
	"vblog/apps/user"
	"vblog/utils/errors"
	"vblog/utils/response"

	"github.com/gin-gonic/gin"
)

// CreateBlog 创建博客：/vblog/api/v1/blogs
// POST 方法
func (h *handler) CreateBlog(ctx *gin.Context) {

	in := blog.NewBody()

	// 这和业务无关的错误
	if err := ctx.BindJSON(in); err != nil {
		response.SendFailed(ctx, errors.NewBadRequestError(err.Error()))
		return
	}

	//  获取用户
	tk, ok := ctx.Get(user.REQUEST_CTX_TOKEN_NAME)
	if !ok {
		response.SendFailed(ctx, errors.NewUnauthorizedError("Not LOGIN, please Login first"))
		return
	}

	in.Author = tk.(*user.Token).Username

	// 调用业务的 CreateBlog
	ins, err := h.svc.CreateBlog(ctx.Request.Context(), in)
	if err != nil {
		// 处理异常
		response.SendFailed(ctx, err)
		return
	}

	response.SendSuccess(ctx, ins)
}

// QueryBlog 查询博客：/vblog/api/v1/blogs
// GET 方法，使用 query 参数
func (h *handler) QueryBlog(ctx *gin.Context) {
	in := blog.NewQueryBlogRequest()
	if err := ctx.BindQuery(in); err != nil {
		response.SendFailed(ctx, errors.NewBadRequestError(err.Error()))
		return
	}

	set, err := h.svc.QueryBlog(ctx, in)
	if err != nil {
		// 处理异常
		response.SendFailed(ctx, err)
		return
	}
	response.SendSuccess(ctx, set)
}

// DescribeBlog 获取博客详情：/vblog/api/v1/blogs/:id
// GET 方法
func (h *handler) DescribeBlog(ctx *gin.Context) {
	in := blog.NewDescribeBlogRequest(0)

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.SendFailed(ctx, errors.NewBadRequestError(err.Error()))
		return
	}
	in.Id = int(id)

	ins, err := h.svc.DescribeBlog(ctx, in)
	if err != nil {
		// 处理异常
		response.SendFailed(ctx, err)
		return
	}
	response.SendSuccess(ctx, ins)
}

// UpdateBlog 更新博客：/vblog/api/v1/blogs
// 使用 PATCH 方法
func (h *handler) UpdateBlog(ctx *gin.Context) {
	in := blog.NewUpdateBlogRequest()

	if err := ctx.BindJSON(in); err != nil {
		response.SendFailed(ctx, err)
		return
	}

	ins, err := h.svc.UpdateBlog(ctx, in)
	if err != nil {
		// 处理异常
		response.SendFailed(ctx, err)
		return
	}
	response.SendSuccess(ctx, ins)
}

// DeleteBlog 删除博客 /vblog/api/v1/blogs/:id
// DELETE 方法
func (h *handler) DeleteBlog(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		response.SendFailed(ctx, errors.NewBadRequestError(err.Error()))
		return
	}

	in := blog.NewDeleteBlogRequest(id)

	ins, err := h.svc.DeleteBlog(ctx, in)
	if err != nil {
		// 处理异常
		return
	}
	response.SendSuccess(ctx, ins)
}
