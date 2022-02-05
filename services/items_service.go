package services

import (
	"errors"

	"context"

	"github.com/ThiyagoNearle/bookstore_users-api/domain/users"
	"github.com/ThiyagoNearle/bookstore_users-api/middleware"
)

func Getuser(ctx context.Context) (*users.User, error) { // it is function not method

	user, _ := middleware.GetCustomContext(ctx)

	print("usid==", user.Id)

	if user.Id == 0 {
		return nil, errors.New("cant able to id information")
	} else {
		return user, nil
	}
}
