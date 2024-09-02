package impl

import "vblog/apps/comment"

var _ comment.Service = &Impl{}

type Impl struct {
}
