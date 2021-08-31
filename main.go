//©2021 Matthew R Kasun mkasun@nusak.ca

package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

var Data PageData

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	router := SetupRouter()
	router.Run(":8080")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("html/*")
	router.Static("images", "./images")
	router.StaticFile("favicon.ico", "./images/favicon.ico")
	router.GET("/", DisplayLanding)
	router.POST("/create_network", CreateNetwork)
	return router
}
