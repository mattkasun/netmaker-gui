//©2021 Matthew R Kasun mkasun@nusak.ca

package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/gravitl/netmaker/database"
)

var Data PageData

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	if err := database.InitializeDatabase(); err != nil {
		log.Fatal("Error connecting to Database", err)
	}
	router := SetupRouter()
	router.Run("127.0.0.1:8080")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("html/*")
	router.Static("images", "./images")
	router.StaticFile("favicon.ico", "./images/favicon.ico")
	router.GET("/", DisplayLanding)
	router.POST("/create_network", CreateNetwork)
	router.POST("/edit_network", EditNetwork)
	router.POST("/delete_network", DeleteNetwork)
	router.POST("/update_network", UpdateNetwork)
	return router
}
