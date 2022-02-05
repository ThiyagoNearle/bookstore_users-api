package services

import (
	"fmt"

	"github.com/ThiyagoNearle/bookstore_users-api/domain/users"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/accessToken"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/credentials"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/date_utils"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/errors"
)

var (
	UserService userServiceInterface = &userService{} // UserService is a variable of type interface so it can return any data type values ( here it return struct type values)
	// That interface which is assigned to that variable is consist some methods,
	// so with the help of this variable we can access these methods.
)

type userService struct{}

type userServiceInterface interface { // An interface can be any data type and can hold atleast 1 method
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(user users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
	LoginUser(users.LoginRequest) (*user.users, *errors.RestErr)
}

// assume that we get a valid id
func (s *userService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil { // this Get method not assigned in the interface, so we directly call variable.method()
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

	user.Password = credentials.HashPassword(user.Password)

	if err := user.SaveUser(); err != nil {
		return nil, err
	}
	if err := user.SaveUserProfile(); err != nil {
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

func (s *userService) LoginUser(requestParam users.LoginRequest) (*users.User, *errors.RestErr) {
	var user users.User

	db_users, status := user.CheckUserName(requestParam)
	if status == false {
		return nil, errors.NewsBadRequestError("invalid login credentials")
	}

	if db_users.Id == 0 && db_users.Configid == 0 {
		return nil, errors.NewNotFoundError("user id & config_id must be valid")

	}

	token, err := accessToken.GenerateToken(int(user.Id), user.Configid) // PASSING USER ID & CONFIGID TO A FUNCTION THAT CREATES TOKEN STRING
	if err != nil {
		return nil, errors.UnauthorizedError("error when generating token")
	}
	fmt.Println("token", token)

	result, err := db_users.LoginResponse(int64(db_users.Id))
	if err != nil {
		return nil, errors.NewsBadRequestError("no user profiles with this id")
	}
	return result, nil
}

/*             or go with GetMD5 method
	dao := &users.User{
		Email:    requestParam.Email,
		Password: converted_pass,
	}
	if err := dao.Login(); err != nil {
		return nil, err
	}
	return dao, nil
}

*/
