package cacoo

import (
	"context"
	"fmt"
)

// UsersService handles user related methods
type UsersService service

// User and Account structs use the same data. this is why Account is used as
// User. In the future, if user data changes, then no need to change any line in
// the function, just you can modify the User type.
type User Account

// Get fetches the user data with given parameters
// Example : https://cacoo.com/api/v1/name.format
func (us *UsersService) Get(ctx context.Context, name string) (*User, *Response, error) {
	u := fmt.Sprintf("users/%s.json", name)

	user := new(User)

	resp, err := us.client.Get(ctx, u, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}
