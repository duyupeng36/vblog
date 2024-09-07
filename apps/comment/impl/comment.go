package impl

import (
	"context"
	"vblog/apps/comment"
	"vblog/utils"
	"vblog/utils/errors"
)

func (i *impl) CreateComment(ctx context.Context, body *comment.Body) (*comment.Comment, error) {

	// 校验用户提供的参数
	if err := utils.Validate(body); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	// 校验成功，创建一个 Comment 对象
	ins := comment.NewComment(body)

	// 入库

	if err := i.db.Create(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) QueryComment(ctx context.Context, request *comment.QueryCommentRequest) (*comment.CommentSet, error) {
	return nil, nil
}

func (i *impl) DescribeComment(ctx context.Context, request *comment.DescribeCommentRequest) (*comment.Comment, error) {
	return nil, nil
}

func (i *impl) UpdateComment(ctx context.Context, request *comment.UpdateCommentRequest) (*comment.Comment, error) {
	return nil, nil
}

func (i *impl) DeleteComment(ctx context.Context, request *comment.DeleteCommentRequest) (*comment.Comment, error) {
	return nil, nil
}
