package comment

type CommentSet struct {
	Comments []*Comment
}

type Comment struct {
	ID int `json:"id"`

	UserID  int    `json:"user_id"` // 谁评论的
	BlogID  int    `json:"blog_id"` // 评论的那篇文章
	Content string `json:"content"` // 评论内容
}
