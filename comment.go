package cacoo

import (
	"context"
	"fmt"
)

// CommentService handles comment related methods
type CommentService service

// Comment holds the information about the comment of diagrams
type Comment struct {
	User    *Account   `json:"user,omitempty"`
	Content *string    `json:"content,omitempty"`
	Updated *CacooTime `json:"updated,omitempty"`
	Created *CacooTime `json:"created,omitempty"`
}

// CommentOption holds the options for comment parameters
type CommentOption struct {
	Name string `json:"name,omitempty"`
}

// Post posts a comment for a diagram
func (cs *CommentService) Post(ctx context.Context, diagramID string, opt *CommentOption) (*Comment, *Response, error) {
	u := fmt.Sprintf("diagrams/%s/comments/post.json", diagramID)

	c := new(Comment)

	resp, err := cs.client.Post(ctx, u, opt, &c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}
