package impl_test

import (
	"context"
	"testing"
	"vblog/apps/comment"
)

var (
	ctx = context.Background()
)

func TestCreateComment(t *testing.T) {
	in := comment.NewBody()
	in.UserID = 1
	in.BlogID = 2
	in.Content = "test"
	ins, err := controller.CreateComment(ctx, in)
	if err != nil {
		t.Error(err)
	}

	t.Log(ins)
}
