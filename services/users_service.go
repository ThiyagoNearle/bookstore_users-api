package services

import (
	"github.com/ThiyagoNearle/bookstore_users-api/domain/users"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/errors"
)

// assume that we get a valid id
func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	err := user.Validate()

	if err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
