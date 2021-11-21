package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	controller "github.com/gravitl/netmaker/controllers"
	"github.com/gravitl/netmaker/logic"
	"github.com/gravitl/netmaker/models"
)

//EditNode display a form to update a node
func EditNode(c *gin.Context) {
	network := c.PostForm("network")
	mac := c.PostForm("mac")
	var node models.Node
	node, err := controller.GetNode(mac, network)
	if err != nil {
		fmt.Println("error getting node details \n", err)
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	c.HTML(http.StatusOK, "EditNode", node)
}

//DeleteNode delele node
func DeleteNode(c *gin.Context) {
	mac := c.PostForm("mac")
	net := c.PostForm("net")
	fmt.Println("deleting node ", mac, net)
	err := controller.DeleteNode(mac+"###"+net, false)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	ReturnSuccess(c, "Nodes", "node deleted")
}

//UpdateNode updates a node
func UpdateNode(c *gin.Context) {
	var node *models.Node
	if err := c.ShouldBind(&node); err != nil {
		fmt.Println("should bind")
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	net := c.Param("net")
	mac := c.Param("mac")
	fmt.Printf("=============%T %T %T %v %v %v", net, mac, node, net, mac, node)
	oldnode, err := logic.GetNodeByMacAddress(net, mac)
	if err != nil {
		fmt.Println("Get node with mac ", mac, " and Network ", net)
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	if err = logic.UpdateNode(&oldnode, node); err != nil {
		fmt.Println("update network")
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	ReturnSuccess(c, "Nodes", "node updated")
}

//NodeHealth return the last checkin time including health status
//and color code for all nodes
func NodeHealth(c *gin.Context) {
	nodes, err := logic.GetAllNodes()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var response []NodeStatus
	var nodeHealth NodeStatus
	for _, node := range nodes {
		nodeHealth.Mac = node.MacAddress
		nodeHealth.Network = node.Network
		lastupdate := time.Now().Sub(time.Unix(node.LastCheckIn, 0))
		if lastupdate.Minutes() > 15.0 {
			nodeHealth.Status = "Error: Node last checked in more than 15 minutes ago: "
			nodeHealth.Color = "w3-deep-orange"
		} else if lastupdate.Minutes() > 5.0 {
			nodeHealth.Status = "Warning: Node last checked in more than 5 minutes ago: "
			nodeHealth.Color = "w3-khaki"
		} else {
			nodeHealth.Status = "Healthy: Node checked in within the last 5 minutes: "
			nodeHealth.Color = "w3-teal"
		}
		response = append(response, nodeHealth)
	}
	c.JSON(http.StatusOK, response)
	return
}
