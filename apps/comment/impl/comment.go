package impl

import (
	"context"
	"net/http"
	"vblog/apps/comment"
	"vblog/utils"
)

const (
	StatusOK          = 200 // 正常
	StatusParamsError = 300 // 请求参数错误
	StatusStoreError  = 400 // 入库失败
)

func (i *Impl) CreateComment(ctx context.Context, body *comment.Body) (*comment.Comment, error) {

	// 校验用户提供的参数
	if err := utils.Validate(body); err != nil {
		return nil, utils.NewAPIError(StatusParamsError, err.Error()).SetHttpStatus(http.StatusBadRequest)
	}

	// 校验成功，创建一个 Comment 对象
	ins := comment.NewComment(body)

	// 入库

	if err := i.db.Create(ins).Error; err != nil {
		return nil, utils.NewAPIError(StatusStoreError, err.Error()).SetHttpStatus(http.StatusInternalServerError)
	}

	return ins, nil
}

func (i *Impl) QueryComment(ctx context.Context, request *comment.QueryCommentRequest) (*comment.CommentSet, error) {
	return nil, nil
}

func (i *Impl) DescribeComment(ctx context.Context, request *comment.DescribeCommentRequest) (*comment.Comment, error) {
	return nil, nil
}

func (i *Impl) UpdateComment(ctx context.Context, request *comment.UpdateCommentRequest) (*comment.Comment, error) {
	return nil, nil
}

func (i *Impl) DeleteComment(ctx context.Context, request *comment.DeleteCommentRequest) (*comment.Comment, error) {
	return nil, nil
}
