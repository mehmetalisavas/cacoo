package cacoo

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

const TestToken = "123456789qwert"

func TestMyAccountInformation(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account.json", func(w http.ResponseWriter, r *http.Request) {
		method(t, r, "GET")
		fmt.Fprint(w, `{
	    "name": "1MUJPfNEEeVUox15",
	    "nickname": "Yoko",
	    "type": "cacoo",
    	"imageUrl": "https://cacoo.com/account/1MUJPfNEEeVUox15/image/32x32"
		}`)
	})

	account, _, err := client.Account.MyAccountInformation(context.Background())
	if err != nil {
		t.Errorf("account.myaccountinformation has error: %v", err)
	}

	want := testAccount()

	if !reflect.DeepEqual(account, want) {
		t.Errorf("want: %v, but got: %v", want, account)
	}
}
