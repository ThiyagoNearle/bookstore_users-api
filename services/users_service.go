package services

import (
	"github.com/ThiyagoNearle/bookstore_users-api/domain/users"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/crypto_utils"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/date_utils"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/errors"
)

var (
	UserService userServiceInterface = &userService{} // holding struct value
)

type userService struct{}

type userServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(user users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
}

// assume that we get a valid id
func (s *userService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	err := user.Validate()

	if err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()

	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil { // err := &RestErr{values}          /// here we passing values to methods ( so the methods in doamin that design as receiver as pointer in domain.)
		/// so if we pass values , then the domain function will take the address
		return nil, err
	}
	return &user, nil
}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current := &users.User{Id: user.Id} //current, err:= users.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}

	//if err := user.Validate(); err != nil {
	//	return nil, err
	//}

	if isPartial { // isPartial means isPartial == True go to this loop
		if user.FirstName != "" { // if we didn'y give any value that consider as empty string
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.LastName = user.Email
		}

	} else {

		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *userService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()

}

func (s *userService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)

}
