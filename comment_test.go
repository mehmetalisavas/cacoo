package cacoo

import (
	"context"
	"fmt"
	"testing"

	"github.com/bookerzzz/grok"
)

func TestCommentPost(t *testing.T) {
	c := NewClient(TestToken)
	// c := NewClient("asd")

	opt := &CommentOption{
		Name: "Comment information - test - cacoo",
	}

	// comment, _, err := c.Comment.Post(context.Background(), "rkzgn5jGe2asBRCQ", opt)
	comment, _, err := c.Comment.Post(context.Background(), "dFevldI74IS92lpr", opt)
	if err != nil {
		fmt.Println("err is:", err)
	}

	grok.V(comment)
	fmt.Println(comment)
	_ = comment
}
