package users

import (
	"strings"

	"github.com/ThiyagoNearle/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
	Configid    int    `json:"configid"`
	Age         string `json:"age"`
	Shopname    string `json:"shopname"`
	Location    string `json:"location"`
}

type User_profile struct {
	Id       int64  `json:"id"`
	Age      string `json:"age"`
	Shopname string `json:"shopname"`
	Location string `json:"location"`
}

type Display_result struct {
	User         User         `json:"user"`
	User_profile User_profile `json:"user_profile"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName) // TrimSpace remove front & back space of the value in json body if we give value like empty "   ", (there is no value) & trimspace will give "" because there is no character.
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
