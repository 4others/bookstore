package app

import "github.com/gin-gonic/gin"

var router = gin.Default()

//StartApplication is a function responsible for running server
func StartApplication() {
	mapUrls()
	router.Run(":8082")
}
