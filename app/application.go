package app

import (
	"github.com/ThiyagoNearle/bookstore_users-api/logger"
	"github.com/ThiyagoNearle/bookstore_users-api/middleware"
	"github.com/gin-gonic/gin"
)

var router = gin.Default() // gin.Default() returns a *engine , router is private variable so it valid only inside the app folder(within a folder if you created many files,that means all files are same like 1 file)

func StartApplication() {
	logger.Info("about to start the application...")
	contextkey := "nearle"
	router.Use(middleware.CustomTokenMiddleware(contextkey))
	mapUrls()
	router.Run(":8080")

}
