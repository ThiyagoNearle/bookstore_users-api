package app

import (
	"github.com/ThiyagoNearle/bookstore_users-api/controllers/ping"
	"github.com/ThiyagoNearle/bookstore_users-api/controllers/users"
)

func mapUrls() {

	router.GET("/ping", ping.Ping) // sub folder name where the file is located.function name

	router.GET("/users/:user_id", users.GetUser) // name of the domain.action that we performed on the domain
	router.POST("/user", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
	router.POST("/Login/user", users.Login)

}

// 	_ "github.com/ThiyagoNearle/bookstore_users_api/controllers/ping"
//	_ "github.com/ThiyagoNearle/bookstore_users_api/controllers/users"
