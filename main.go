//©2021 Matthew R Kasun mkasun@nusak.ca

package main

import (
	"github.com/gin-gonic/gin"
)

var Data PageData

func main() {
	router := SetupRouter()
	router.Run(":8080")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("html/*")
	router.GET("/", DisplayLanding)
	return router
}
