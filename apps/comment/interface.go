package comment

import "context"

const (
	AppName = "comments"
)

// 博客管理的业务接口声明

type Service interface {
	CreateComment(ctx context.Context, body *Body) (*Comment, error)
	QueryComment(ctx context.Context, request *QueryCommentRequest) (*CommentSet, error)
	DescribeComment(ctx context.Context, request *DescribeCommentRequest) (*Comment, error)
	UpdateComment(ctx context.Context, request *UpdateCommentRequest) (*Comment, error)
	DeleteComment(ctx context.Context, request *DeleteCommentRequest) (*Comment, error)
}

type QueryCommentRequest struct {
}

type DescribeCommentRequest struct{}

type UpdateCommentRequest struct{}

type DeleteCommentRequest struct{}
