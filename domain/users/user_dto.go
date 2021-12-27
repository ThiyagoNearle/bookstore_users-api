package users

import (
	"strings"

	"github.com/ThiyagoNearle/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct { // call always users.User
	Id          int64  `json:"id"`
	FirstNanme  string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"-"` // this field, internally we working with password & dont want this as a JSON
}

func (user *User) Validate() *errors.RestErr {
	user.FirstNanme = strings.TrimSpace(user.FirstNanme) // in json body if we give value like empty "   ", (there is no value) & trimspace will give "" because there is no character.
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" { // if we didn't give email address also it treats as empty string like ""
		return errors.NewsBadRequestError("invalid email address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewsBadRequestError("invalid password")
	}
	return nil

}
