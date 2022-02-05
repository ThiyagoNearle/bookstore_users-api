package users

import (
	"net/http"
	"strconv"

	"github.com/ThiyagoNearle/bookstore_users-api/domain/users"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/errors"

	"github.com/ThiyagoNearle/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	UserId, userErr := strconv.ParseInt(userIdParam, 10, 64) // this function return 2 valuesand you save those values in these 2 objects variables

	if userErr != nil {
		return 0, errors.NewsBadRequestError("user id should be a number")
	}
	return UserId, nil
}

func Create(c *gin.Context) {
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

	result, saveErr := services.UserService.CreateUser(user) // attempt to create this user in the database
	// if we are getting 404 error from the services then this controller is just returning that error
	// if we have any error on the services, we can just return that error as json

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr) // though saveErr hold address of error value, if we send that to json, it takes only value out of that
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	// when we run the app intially we store contextkey and user in that context.
	// c.Request.Context() => getting the context(what ever may be in the context) from the request
	result, err := services.Getuser(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewNotFoundError("error in data"))
		return
	}
	c.JSON(http.StatusOK, result)
	return
}

func Update(c *gin.Context) {
	//-------------------------------------------------GET(from URL)----------------------------------------------	-

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	//---------------------------------------------------PUT(from JSON BODY )--------------------------------

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewsBadRequestError(("invalid json body"))
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId /// copy id to

	isPartial := (c.Request.Method == http.MethodPatch) // it gives true if the method is PATCH

	result, err := services.UserService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)

}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UserService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UserService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))

	// while rendering as we display json if we want to call some "METHOD" in some other package, we dont need to call that package name instead we can directly call the method with our variable.

}

func Login(c *gin.Context) {
	var request users.LoginRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		bodyErr := errors.NewsBadRequestError("invalid json body")
		c.JSON(bodyErr.Status, bodyErr)
		return
	}

	final_result, saveErr := services.UserService.LoginUser(request) //LoginUser(request)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, final_result)
}
