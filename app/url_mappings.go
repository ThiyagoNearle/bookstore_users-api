package app

import (
	"github.com/ThiyagoNearle/bookstore_users-api/controllers/ping"
	"github.com/ThiyagoNearle/bookstore_users-api/controllers/users"
)

func mapUrls() {

	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/user", users.CreateUser)

}

// 	_ "github.com/ThiyagoNearle/bookstore_users_api/controllers/ping"
//	_ "github.com/ThiyagoNearle/bookstore_users_api/controllers/users"
