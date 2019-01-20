package cacoo

import (
	"context"
	"fmt"
)

// AccountService handles account related methods
type AccountService service

// Account represents the information about the account on cacoo
type Account struct {
	ImageURL *string `json:"imageUrl,omitempty"`
	Name     *string `json:"name,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
	Type     *string `json:"type,omitempty"`
}

// MyAccountInformation fetches the information about the current user
func (as *AccountService) MyAccountInformation(ctx context.Context) (*Account, *Response, error) {
	u := fmt.Sprintf("%s", "account.json")

	a := new(Account)

	resp, err := as.client.Get(ctx, u, &a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}
