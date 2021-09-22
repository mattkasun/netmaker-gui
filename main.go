//©2021 Matthew R Kasun mkasun@nusak.ca

package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	controller "github.com/gravitl/netmaker/controllers"
	"github.com/gravitl/netmaker/database"
	"github.com/gravitl/netmaker/servercfg"
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
	fmt.Println("using database: ", servercfg.GetDB())
	if err := database.InitializeDatabase(); err != nil {
		log.Fatal("Error connecting to Database:\n", err)
	}
	router := gin.Default()
	store := memstore.NewStore([]byte("secret"))

	router.Use(sessions.Sessions("netmaker", store))
	router.LoadHTMLGlob("html/*")
	router.Static("images", "./images")
	router.StaticFile("favicon.ico", "./images/favicon.png")
	router.POST("/newuser", NewUser)
	router.POST("/login", ProcessLogin)
	//use  authorization middleware
	private := router.Group("/", AuthRequired)
	{
		//router.Use(AuthRequired)
		private.GET("/", DisplayLanding)
		private.POST("/create_network", CreateNetwork)
		private.POST("/edit_network", EditNetwork)
		private.POST("/delete_network", DeleteNetwork)
		private.POST("/update_network", UpdateNetwork)
		private.GET("/refreshkeys/:net", RefreshKeys)
		private.POST("/create_key", NewKey)
		private.POST("/delete_key", DeleteKey)
		private.POST("/create_user", CreateUser)
		private.POST("/delete_user", DeleteUser)
		private.POST("/edit_node", EditNode)
		private.POST("/delete_node", DeleteNode)
		private.POST("/update_node/:net/:mac", UpdateNode)
		private.GET("/node_health", NodeHealth)
		private.POST("/create_egress/:net/:mac", CreateEgress)
		private.POST("/process_egress/:net/:mac", ProcessEgress)
		private.POST("/delete_egress/:net/:mac", DeleteEgress)
		private.POST("/create_ingress/:net/:mac", CreateIngress)
		private.POST("/delete_ingress/:net/:mac", DeleteIngress)
		private.POST("/create_ingress_client/:net/:mac", CreateIngressClient)
		private.POST("/delete_ingress_client/:net/:id", DeleteIngressClient)
		private.POST("/edit_ingress_client/:net/:id", EditIngressClient)
		private.POST("/get_qr/:net/:id", GetQR)
		private.POST("/get_client_config/:net/:id", GetClientConfig)
		private.POST("/update_client/:net/:id", UpdateClient)

		private.GET("/logout", LogOut)
	}
	return router
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	options := session.Options

	fmt.Println("checking authorization\n", options)
	fmt.Printf("type %v value %s\n", options, options)
	loggedIn := session.Get("loggedIn")
	fmt.Println("loggedIn status: ", loggedIn)
	if loggedIn != true {
		adminExists, err := controller.HasAdmin()
		fmt.Println("response from HasAdmin(): ", adminExists, err)
		if err != nil {
			fmt.Println("error checking for admin")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
		}
		if !adminExists {
			fmt.Println("no admin")
			c.HTML(http.StatusOK, "new", nil)
			c.Abort()
		} else {
			message := session.Get("error")
			fmt.Println("user exists --- message\n", message)
			c.HTML(http.StatusUnauthorized, "Login", gin.H{"messge": message})
			c.Abort()
		}
	}
	fmt.Println("authorized - good to go")
}
