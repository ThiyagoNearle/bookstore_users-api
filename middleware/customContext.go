package middleware

import (
	"context"
	"errors"
	"fmt"

	"github.com/ThiyagoNearle/bookstore_users-api/domain/users"
)

func GetCustomContext(ctx context.Context) (*users.User, error) {

	var userCtxKey string
	userCtxKey = "userCtxKey"

	if ctx.Value(userCtxKey) == nil { // context holding some values(ex: user), inorder to get the user, in context.Value, we need to pass the same key

		return nil, errors.New("contextkey null") // means the context key is empty....
	}
	user, ok := ctx.Value(userCtxKey).(*users.User) // for context output type like this only moreooover in coontext we have user = &data
	fmt.Println("----------------------------------")
	fmt.Printf("%+v", user)
	fmt.Println("----------------------------------")
	if !ok || user.Id == 0 {

		return nil, errors.New("contextkey null") // or no record in the context
	}
	// if everything goes correct then we get user
	return user, nil
}
