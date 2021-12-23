package users

import (
	"fmt"

	"github.com/ThiyagoNearle/bookstore_users-api/datasources/mysql/users_db"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/date_utils"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/errors"
)

var usersDB = make(map[int64]*User)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstNanme = result.FirstNanme
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewsBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewsBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.DateCreated = date_utils.GetNowString()

	usersDB[user.Id] = user // current = user
	return nil
}

//ery4y4y4y
