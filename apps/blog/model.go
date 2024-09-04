package blog

import (
	"encoding/json"
	"time"
	"vblog/utils"
)

type BlogSet struct {
	Total int64   `json:"total"`
	Items []*Blog `json:"items"`
}

func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

func (set *BlogSet) String() string {
	b, _ := json.MarshalIndent(set, "", "  ")
	return string(b)
}

// Blog 代表文章的数据结构
type Blog struct {
	*Meta
	*Body
}

func NewBlog(body *Body) *Blog {
	return &Blog{
		Meta: NewMeta(),
		Body: body,
	}
}

func (b *Blog) String() string {

	if result, err := json.MarshalIndent(b, "", "  "); err != nil {
		return ""
	} else {
		return string(result)
	}
}

// Meta 文章元数据
type Meta struct {
	ID          int   `json:"id"`           // 文章 ID 由数据库生成
	CreatedAt   int64 `json:"created_at"`   // 创建时间 由系统生成
	UpdatedAt   int64 `json:"updated_at"`   // 更新时间 由系统生成
	PublishedAt int64 `json:"published_at"` // 发布时间 由系统生产
}

func NewMeta() *Meta {
	return &Meta{
		CreatedAt: time.Now().Unix(),
	}
}

// Body 用户提供的文章数据
type Body struct {
	Title   string            `json:"title" validate:"required"`
	Author  string            `json:"Author" validate:"required"`
	Content string            `json:"content" validate:"required"`
	Tags    map[string]string `json:"tags" gorm:"serializer:json"`
	Status  STATUS            `json:"status"`
}

func NewBody() *Body {
	return &Body{
		Tags:   make(map[string]string),
		Status: DRAFT,
	}
}

func (b *Body) Validate() error {

	// 验证请求参数是否正确提供
	// 我可以执行编写验证逻辑
	// if b.Content == "" {
	// 	return fmt.Errorf("content must be required")
	// }

	// 验证请求参数是每个业务都可能需要的。因此，我们需要将验证逻辑独立到 utils 包中，方便所有的业务使用
	return utils.Validate(b)
}
