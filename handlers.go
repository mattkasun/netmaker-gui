package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gravitl/netmaker/models"
)

func DisplayLanding(c *gin.Context) {
	var Data PageData
	Data.Init("Networks")
	c.HTML(http.StatusOK, "layout", Data)
}

func CreateNetwork(c *gin.Context) {
	var net models.Network

	net.NetID = c.PostForm("name")
	net.AddressRange = c.PostForm("address")
	net.IsDualStack = c.PostForm("dual")
	net.IsLocal = c.PostForm("local")
	net.DefaultUDPHolePunch = c.PostForm("udp")

	response, err := API(net, http.MethodPost, "/api/networks", "secretkey")
	fmt.Println(err, response)
	var message models.ErrorResponse
	json.NewDecoder(response.Body).Decode(&message)
	fmt.Println(message)
	if err != nil {

		c.JSON(http.StatusBadRequest, response)
	}
	location := url.URL{Path: "/"}
	c.Redirect(http.StatusFound, location.RequestURI())

}
