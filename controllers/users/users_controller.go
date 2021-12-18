package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ThiyagoNearle/bookstore_users-api/domain/dusers"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user dusers.User // users is a package under domain package & User is a struct
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// TODO: Handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		// TODO: Handle error
		return
	}
	fmt.Println(user)
	c.String(http.StatusNotModified, "implement me")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotModified, "implement me")
}
