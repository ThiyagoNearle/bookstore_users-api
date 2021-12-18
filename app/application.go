package app

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default() // router is private variable so it valid only inside the app folder(within a folder if you created many files,that means all files are same like 1 file)

func StartApplication() {
	mapUrls() // we can use because this function under same folder only but another file ( so act as a same file)
	router.Run(":8080")

}
