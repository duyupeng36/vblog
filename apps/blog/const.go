package blog

type STATUS uint8

const (
	// DRAFT : 文章的默认状态，草稿状态
	DRAFT STATUS = iota
	UNPUBLISHED
	PUBLISHED
)
