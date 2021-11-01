package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	controller "github.com/gravitl/netmaker/controllers"
	"github.com/gravitl/netmaker/logic"
	"github.com/gravitl/netmaker/models"
)

// ProcessEgress adds the node as an Egress Gateway
func ProcessEgress(c *gin.Context) {
	var egress models.EgressGatewayRequest
	egress.NodeID = c.Param("mac")
	egress.NetID = c.Param("net")
	egress.Ranges = strings.Split(c.PostForm("ranges"), ",")
	egress.Interface = c.PostForm("interface")
	node, err := controller.CreateEgressGateway(egress)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	session := sessions.Default(c)
	session.Set("message", node.Name+" is now an egress gateway")
	session.Set("page", "Nodes")
	session.Save()
	location := url.URL{Path: "/"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

// CreateEggress displays modal for Egress Gateway Creation
func CreateEgress(c *gin.Context) {
	net := c.Param("net")
	mac := c.Param("mac")
	node, err := controller.GetNode(mac, net)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	c.HTML(http.StatusOK, "Egress", node)
}

// DeleteEgress removes the node as an Egress Gateway
func DeleteEgress(c *gin.Context) {
	net := c.Param("net")
	mac := c.Param("mac")
	_, err := controller.DeleteEgressGateway(net, mac)
	if err != nil {
		fmt.Println(err)
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	session := sessions.Default(c)
	session.Set("message", "Egress Gateway has been deleted")
	session.Set("page", "Nodes")
	session.Save()
	location := url.URL{Path: "/"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

// CreateIngress add the node as an Ingress Gateway
func CreateIngress(c *gin.Context) {
	net := c.Param("net")
	mac := c.Param("mac")
	node, err := controller.CreateIngressGateway(net, mac)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	session := sessions.Default(c)
	session.Set("message", node.Name+" is now an ingress gateway")
	session.Set("page", "Nodes")
	session.Save()
	location := url.URL{Path: "/"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func DeleteIngress(c *gin.Context) {
	net := c.Param("net")
	mac := c.Param("mac")
	_, err := controller.DeleteIngressGateway(net, mac)
	if err != nil {
		fmt.Println(err)
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ingress Gateway Deleted"})
}

func CreateRelay(c *gin.Context) {
	var relayData Relay
	mac := c.Param("mac")
	net := c.Param("net")
	node, err := controller.GetNode(mac, net)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	nodes, err := logic.GetNetworkNodes(node.Network)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	relayData.Node = node
	relayData.Nodes = nodes
	c.HTML(http.StatusOK, "CreateRelay", relayData)
}

func ProcessRelayCreation(c *gin.Context) {
	var request models.RelayRequest
	request.NodeID = c.Param("mac")
	request.NetID = c.Param("net")
	request.RelayAddrs = c.PostFormArray("address")
	_, err := controller.CreateRelay(request)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	session := sessions.Default(c)
	session.Set("message", "Relay Gateway Created")
	session.Set("page", "Nodes")
	session.Save()
	location := url.URL{Path: "/"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func DeleteRelay(c *gin.Context) {
	net := c.Param("net")
	mac := c.Param("mac")
	_, err := controller.DeleteRelay(net, mac)
	if err != nil {
		fmt.Println(err)
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	session := sessions.Default(c)
	session.Set("message", "Relay Gateway Deleted")
	session.Set("page", "Nodes")
	session.Save()
	location := url.URL{Path: "/"}
	c.Redirect(http.StatusFound, location.RequestURI())
}
