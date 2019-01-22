package cacoo

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCommentPost(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/diagrams/1MUJPfNEEeVUox15/comments/post.json", func(w http.ResponseWriter, r *http.Request) {
		method(t, r, "POST")
		fmt.Fprint(w, `{
		    "user": {
		        "name": "1MUJPfNEEeVUox15",
		        "nickname": "Yoko",
		        "type": "cacoo",
		        "imageUrl": "https://cacoo.com/account/1MUJPfNEEeVUox15/image/32x32"
		    },
		    "content": "comment 1",
		    "created": "Mon, 10 Aug 2009 17:00:00 +0900",
		    "updated": "Mon, 10 Aug 2009 17:00:00 +0900"
		}`)
	})

	opt := &CommentOption{
		Name: "Comment information",
	}

	comment, _, err := client.Comment.Post(context.Background(), "1MUJPfNEEeVUox15", opt)
	if err != nil {
		t.Errorf("comment.post has error: %v", err)
	}

	c := testComment(t)
	want := &c

	if !reflect.DeepEqual(comment, want) {
		t.Errorf("want: %v, but got: %v", want, comment)
	}
}
