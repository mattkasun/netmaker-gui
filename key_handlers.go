package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	controller "github.com/gravitl/netmaker/controllers"
	"github.com/gravitl/netmaker/models"
)

//RefreshKeys refreshs keys for network
func RefreshKeys(c *gin.Context) {
	net := c.Param("net")
	_, err := controller.KeyUpdate(net)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, err, "Keys")
		return
	}
	ReturnSuccess(c, "Networks", "Keys for "+net+" network updated")
}

//NewKey will generate a new key for a network
func NewKey(c *gin.Context) {
	var key models.AccessKey
	var err error
	net := c.PostForm("network")
	key.Name = c.PostForm("name")
	key.Uses, err = strconv.Atoi(c.PostForm("uses"))
	if err != nil {
		key.Uses = 1
	}
	network, err := controller.GetNetwork(net)
	if err != nil {
		fmt.Println("error retrieving network ", err)
		ReturnError(c, http.StatusBadRequest, err, "Keys")
		return
	}
	_, err = controller.CreateAccessKey(key, network)
	if err != nil {
		fmt.Println("error creating key", err)
		ReturnError(c, http.StatusBadRequest, err, "Keys")
		return
	}
	ReturnSuccess(c, "Keys", "")
}

//DeleteKey delete network keys
func DeleteKey(c *gin.Context) {
	net := c.PostForm("net")
	name := c.PostForm("key")
	fmt.Println("Delete Key params: ", net, name)
	if err := controller.DeleteKey(name, net); err != nil {
		fmt.Println("error deleting key", err)
		ReturnError(c, http.StatusBadRequest, err, "Keys")
		return
	}
	ReturnSuccess(c, "Keys", "")
}
