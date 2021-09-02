package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	controller "github.com/gravitl/netmaker/controllers"
	"github.com/gravitl/netmaker/functions"
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

	err := controller.CreateNetwork(net)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		c.Abort()
	}
	location := url.URL{Path: "/"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func EditNetwork(c *gin.Context) {
	network := c.PostForm("network")
	fmt.Println("editing network ", network)
	var Data models.Network
	Data, err := controller.GetNetwork(network)
	if err != nil {
		fmt.Println("error getting net details \n", err)
		c.JSON(http.StatusBadRequest, err)
	}
	c.HTML(http.StatusOK, "EditNet", Data)
}

func DeleteNetwork(c *gin.Context) {
	network := c.PostForm("network")
	err := controller.DeleteNetwork(network)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		c.Abort()
	}
	location := url.URL{Path: "/"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func UpdateNetwork(c *gin.Context) {
	var network models.Network
	net := c.PostForm("NetID")
	network.NetID = c.PostForm("NetID")
	network.AddressRange = c.PostForm("Address")
	network.LocalRange = c.PostForm("Local")
	network.DisplayName = c.PostForm("Name")
	network.DefaultInterface = c.PostForm("Interface")
	port, err := strconv.Atoi(c.PostForm("Port"))
	if err != nil {
		fmt.Println("error converting port", err)
	}
	network.DefaultListenPort = int32(port)
	network.DefaultPostUp = c.PostForm("PostUp")
	network.DefaultPostDown = c.PostForm("PostDown")
	keep, err := strconv.Atoi(c.PostForm("Keepalive"))
	if err != nil {
		fmt.Println("error converting keepalive", err)
	}
	network.DefaultKeepalive = int32(keep)
	check, err := strconv.Atoi(c.PostForm("CheckinInterval"))
	if err != nil {
		fmt.Println("error converting check interval", err)
	}
	network.DefaultCheckInInterval = int32(check)
	network.IsDualStack = c.PostForm("DualStack")
	if network.IsDualStack == "" {
		network.IsDualStack = "no"
	}
	network.DefaultSaveConfig = c.PostForm("DefaultSaveConfig")
	if network.DefaultSaveConfig == "" {
		network.DefaultSaveConfig = "no"
	}
	network.DefaultUDPHolePunch = c.PostForm("UDPHolePunching")
	if network.DefaultUDPHolePunch == "" {
		network.DefaultUDPHolePunch = "no"
	}
	network.AllowManualSignUp = c.PostForm("AllowManualSignup")
	if network.AllowManualSignUp == "" {
		network.AllowManualSignUp = "no"
	}

	if network.LocalRange == "" {
		network.IsLocal = "no"
	} else {
		network.IsLocal = "yes"
	}
	if network.AddressRange6 == "" {
		network.IsIPv6 = "no"
	} else {
		network.IsIPv6 = "yes"
	}
	if network.AddressRange == "" {
		network.IsIPv4 = "no"
	} else {
		network.IsIPv4 = "yes"
	}
	network.IsGRPCHub = "no"

	oldnetwork, err := controller.GetNetwork(net)
	if err != nil {
		fmt.Println("error getting network ", err)
		c.JSON(http.StatusBadRequest, err)
		c.Abort()
	}
	updaterange, updatelocal, err := oldnetwork.Update(&network)
	if err != nil {
		fmt.Println("error updating network ", err)
		c.JSON(http.StatusBadRequest, err)
		c.Abort()
	}
	if updaterange {
		err = functions.UpdateNetworkNodeAddresses(network.NetID)
		if err != nil {
			fmt.Println("error updating network Node Addresses", err)
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
	}
	if updatelocal {
		err = functions.UpdateNetworkLocalAddresses(network.NetID)
		if err != nil {
			fmt.Println("error updating network Local Addresses", err)
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}
	}

	location := url.URL{Path: "/"}
	c.Redirect(http.StatusFound, location.RequestURI())
}
