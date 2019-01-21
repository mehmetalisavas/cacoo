package cacoo

import (
	"context"
	"fmt"
	"net/http"
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
		t.Errorf("Users.Get returned error: %v", err)
	}

	want := &Account{
		ImageURL: String("https://cacoo.com/account/1MUJPfNEEeVUox15/image/32x32"),
		Name:     String("1MUJPfNEEeVUox15"),
		Nickname: String("Yoko"),
		Type:     String("cacoo"),
	}
	if *user.ImageURL != *want.ImageURL {
		t.Errorf("expected: %v, but got : %v", *want.ImageURL, *user.ImageURL)
	}
	if *user.Name != *want.Name {
		t.Errorf("expected: %v, but got : %v", *want.Name, *user.Name)
	}
	if *user.Nickname != *want.Nickname {
		t.Errorf("expected: %v, but got : %v", *want.Nickname, *user.Nickname)
	}
	if *user.Type != *want.Type {
		t.Errorf("expected: %v, but got : %v", *want.Type, *user.Type)
	}
}
