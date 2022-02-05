package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ThiyagoNearle/bookstore_users-api/domain/users"
	"github.com/ThiyagoNearle/bookstore_users-api/utils/accessToken"
	"github.com/gin-gonic/gin"
)

type CustomResult struct {
	Status        bool      `json:"status"`
	Code          int       `json:"code"`
	Message       string    `json:"message"`
	Expirytime    time.Time `json:"expirytime"`
	Sessionstatus bool      `json:"sessionstatus"`
}

func CustomTokenMiddleware(contextkey string) gin.HandlerFunc {

	fmt.Println("-------------------------------welcome to gin middleware")
	contextkey = "nearle"
	return func(c *gin.Context) {

		token := c.Request.Header.Get("token") // token = "abc123" // "token" should be same in the header name token

		print(token) // "abc123"

		if token == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, "token is empty")
			log.Fatal("token is empty")
			return
		}

		userId, configid, err := accessToken.ParseToken(token)
		if err != nil {

			c.JSON(http.StatusUnauthorized, "Token Denied")
			c.Abort()
			fmt.Println("-----------------------------step4.10")
			return
		}

		id := int(userId)
		id1 := int(configid)

		if id1 == 1 {
			print("configid==1")
			data1 := &users.User{}
			user, status, errrr := data1.UserAuthentication(int64(id)) // user receive *values only
			print(status)
			if errrr != nil {
				c.JSON(http.StatusBadRequest, "user not found")
				c.Abort()
				return
			}

			ctx := context.WithValue(c.Request.Context(), contextkey, user)

			c.Request = c.Request.WithContext(ctx)

			c.Next()

		} else {
			print("configid  mismatched")
			c.JSON(http.StatusUnauthorized, "config mismatched")
			c.Abort()
			return
		}

	}

}
