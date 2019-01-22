package cacoo

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSpecificUser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1MUJPfNEEeVUox15.json", func(w http.ResponseWriter, r *http.Request) {
		method(t, r, "GET")
		fmt.Fprint(w, `{
	    "name": "1MUJPfNEEeVUox15",
	    "nickname": "Yoko",
	    "type": "cacoo",
    	"imageUrl": "https://cacoo.com/account/1MUJPfNEEeVUox15/image/32x32"
		}`)
	})

	user, _, err := client.Users.Get(context.Background(), "1MUJPfNEEeVUox15")
	if err != nil {
		t.Errorf("user.get returned error: %v", err)
	}

	want := &User{
		ImageURL: String("https://cacoo.com/account/1MUJPfNEEeVUox15/image/32x32"),
		Name:     String("1MUJPfNEEeVUox15"),
		Nickname: String("Yoko"),
		Type:     String("cacoo"),
	}

	if !reflect.DeepEqual(user, want) {
		t.Errorf("want: %v, but got: %v", want, user)
	}
}

func TestUserGetEmptyUser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1MUJPfNEEeVUox15.json", func(w http.ResponseWriter, r *http.Request) {
		method(t, r, "GET")
		fmt.Fprint(w, `{
	    "name": "1MUJPfNEEeVUox15",
	    "nickname": "Yoko",
	    "type": "cacoo",
    	"imageUrl": "https://cacoo.com/account/1MUJPfNEEeVUox15/image/32x32"
		}`)
	})

	user, _, err := client.Users.Get(context.Background(), "")
	if err == nil {
		t.Errorf("error should not be nil, but got: %v", err)
	}

	if user != nil {
		t.Errorf("user should be nil, but got: %v", user)
	}

}
