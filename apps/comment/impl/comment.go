package impl

import (
	"context"
	"vblog/apps/comment"
)

func (i *Impl) CreateComment(ctx context.Context, comment *comment.Comment) (*comment.Comment, error) {
	return nil, nil
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
