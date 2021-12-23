package users

import (
	"net/http"
	"strconv"

	"github.com/ThiyagoNearle/bookstore_users-api/domain/users"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/errors"

	"github.com/ThiyagoNearle/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User // users is a package under domain package & User is a struct

	// OPTION 1 ( HANDLING JSON REQUEST)

	//	bytes, err := ioutil.ReadAll(c.Request.Body) // c.Request.body contains what the client given data in jason & readll read that json body data and store it as bytes
	//	if err != nil {
	// TODO: Handle error
	//		return
	//	}
	//	if err := json.Unmarshal(bytes, &user); err != nil { // unmarshal converts bytes into string( like original body data that has been given by request) & store it in to a address
	//		fmt.Println(err.Error())
	// TODO: Handle error
	//		return

	// OPTION 2 ( HANDLING JSON REQUEST)

	err := c.ShouldBindJSON(&user) // ShouldBindJSON do all the operation like get the json body request from client and it holds the result in byte and convert this into string ( original json) and store this value in a address we provide
	if err != nil {
		restErr := errors.NewsBadRequestError("invalid json boby")

		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user) // attempt to create this user in the database
	// if we are getting 404 error from the services then this controller is just returning that error
	// if we have any error on the services, we can just return that error as json

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr) // though saveErr hold address of error value, if we send that to json, it takes only value out of that
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewsBadRequestError("user id should be number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr) // though saveErr hold address of error value, if we send that to json, it takes only value out of that
		return
	}
	c.JSON(http.StatusCreated, user)
}
