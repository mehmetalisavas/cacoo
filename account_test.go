package cacoo

import (
	"context"
	"fmt"
	"testing"

	"github.com/bookerzzz/grok"
)

const TestToken = "123456789qwert"

func TestAccount(t *testing.T) {
	client := NewClient(TestToken)
	// client := NewClient("")

	ctx := context.Background()
	account, resp, err := client.Account.MyAccountInformation(ctx)
	if err != nil {
		fmt.Printf("Err is %+v:", err)
	}

	grok.V(account)
	fmt.Println("resp is:", resp)
}
