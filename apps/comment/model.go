package comment

import (
	"database/sql"
	"encoding/json"
	"time"
)

type CommentSet struct {
	Comments []*Comment
}

type Comment struct {
	*Meta
	*Body
}

func (c *Comment) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func NewComment(body *Body) *Comment {
	return &Comment{
		Meta: NewMeta(),
		Body: body,
	}
}

type Meta struct {
	ID        int           `json:"id"`
	CreatedAt int64         `json:"created_at"`
	DeletedAt sql.NullInt64 `json:"deleted_at" `
}

func NewMeta() *Meta {
	return &Meta{
		CreatedAt: time.Now().Unix(),
	}
}

type Body struct {
	UserID  int    `json:"user_id"  validate:"required"` // 谁评论的
	BlogID  int    `json:"blog_id"  validate:"required"` // 评论的那篇文章
	Content string `json:"content"  validate:"required"` // 评论内容
}

func NewBody() *Body {
	return &Body{}
}
