package users

import (
	"bitbucket/Nearle/ariane/Domain/users"
	userservice "bitbucket/Nearle/ariane/Services/users"
	"bitbucket/Nearle/ariane/Utils/Errors"
	"errors"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	var errResult Errors.CustomError
	var err error
	if err := c.ShouldBindJSON(&user); err != nil {
		formatErr := Errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, formatErr)
		return
	}

	result, err := userservice.CreateUser(user) // calling services

	if err != nil {
		if err.Error() == "Email Already Exists" {
			errResult.Status = false
			errResult.Code = http.StatusConflict
			errResult.Message = "Email Already Exists"
			c.JSON(http.StatusConflict, errResult)
			return

		} else if err.Error() == "Contactno Already Exists" {
			errResult.Status = false
			errResult.Code = http.StatusConflict
			errResult.Message = "Contactno Already Exists"
			c.JSON(http.StatusConflict, errResult)
			return
		} else {
			c.JSON(http.StatusConflict, "Backend Error")
			return
		}

	}

	c.JSON(http.StatusCreated, result)
	return
}
func GetUser(c *gin.Context) {
	result, err := userservice.Getuser(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("error in data"))
		return
	}
	c.JSON(http.StatusOK, result)
	return
}
func GetAllUsers(c *gin.Context) {
	result, err := userservice.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, result)
	return
}
func UpdateUser(c *gin.Context) {
	var user users.User
	var errResult Errors.CustomError
	var err error
	var res users.CustomResult
	if err := c.ShouldBindJSON(&user); err != nil {
		formatErr := Errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, formatErr)
		return
	}

	status, err := userservice.UpdateUser(user)
	if err != nil {
		if err.Error() == "Email Already Exists" {
			errResult.Status = false
			errResult.Code = http.StatusConflict
			errResult.Message = "Email Already Exists"
			c.JSON(http.StatusConflict, errResult)
			return

		} else if err.Error() == "Contactno Already Exists" {
			errResult.Status = false
			errResult.Code = http.StatusConflict
			errResult.Message = "Contactno Already Exists"
			c.JSON(http.StatusConflict, errResult)
			return
		} else {
			errResult.Status = false
			errResult.Code = http.StatusConflict
			errResult.Message = "Backend Error"
			c.JSON(http.StatusConflict, errResult)
			return
		}

	}
	if status == true {
		res.Status = true
		res.Code = http.StatusCreated
		res.Message = "Profile Update Successfully"
		c.JSON(http.StatusAccepted, res)
		return
	} else {
		res.Status = true
		res.Code = http.StatusCreated
		res.Message = "Profile Not Updated"
		c.JSON(http.StatusBadRequest, res)
		return
	}

}
func UserLogin(c *gin.Context) {
	var user users.User
	var errResult Errors.CustomError
	var Result users.LoginResult
	if err := c.ShouldBindJSON(&user); err != nil {
		formatErr := Errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, formatErr)
		return
	}

	res, stat, er := userservice.Login(user)
	if stat == false {
		errResult.Status = false
		errResult.Code = http.StatusBadRequest
		errResult.Message = "Incorrect Username"
		c.JSON(http.StatusBadRequest, errResult)
		return
	}
	print(er)
	Result.Status = true
	Result.Code = http.StatusOK
	Result.Message = "Success"
	Result.Locatoninfo = res.Locatoninfo
	Result.Tenantinfo = res.Tenantinfo
	Result.Userinfo = res.Userinfo
	c.JSON(http.StatusOK, Result)
	return
}
