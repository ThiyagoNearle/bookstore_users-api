package services

import (
	"github.com/ThiyagoNearle/bookstore_users-api/domain/users"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/date_utils"
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
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()

	if err := user.Save(); err != nil { // err := &RestErr{values}
		return nil, err
	}
	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id) //current, err:= users.User{Id: user.Id}
	if err != nil {
		return nil, err
	}

	//if err := user.Validate(); err != nil {
	//	return nil, err
	//}

	if isPartial { // isPartial means isPartial == True go to this loop
		if user.FirstNanme != "" { // if we didn'y give any value that consider as empty string
			current.FirstNanme = user.FirstNanme
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.LastName = user.Email
		}

	} else {

		current.FirstNanme = user.FirstNanme
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()

}

func SearchUser(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)

}
