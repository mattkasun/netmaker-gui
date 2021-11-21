package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	controller "github.com/gravitl/netmaker/controllers"
	"github.com/gravitl/netmaker/logic"
	"github.com/gravitl/netmaker/models"
)

//CreateDNS add new custom dns entry
func CreateDNS(c *gin.Context) {
	var entry models.DNSEntry
	entry.Network = c.PostForm("network")
	entry.Name = c.PostForm("name")
	entry.Address = c.PostForm("address")
	if err := controller.ValidateDNSCreate(entry); err != nil {
		fmt.Println("validation err dns: ", err)
		ReturnError(c, http.StatusBadRequest, err, "DNS")
		return
	}
	_, err := controller.CreateDNS(entry)
	if err != nil {
		fmt.Println("err dns: ", err)
		ReturnError(c, http.StatusBadRequest, err, "DNS")
		return
	}
	ReturnSuccess(c, "DNS", "DNS Entry for "+entry.Name+" created")
}

//DeleteDNS deletes custom DNS entry
func DeleteDNS(c *gin.Context) {
	network := c.Param("net")
	name := c.Param("name")
	if err := controller.DeleteDNS(name, network); err != nil {
		fmt.Println("err dns delete", err)
		ReturnError(c, http.StatusBadRequest, err, "DNS")
		return
	}
	if err := logic.SetDNS(); err != nil {
		fmt.Println("err set dns", err)
		ReturnError(c, http.StatusBadRequest, err, "DNS")
		return
	}
	ReturnSuccess(c, "DNS", "DNS Entry deleted")
}
