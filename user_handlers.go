package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	controller "github.com/gravitl/netmaker/controllers"
	"github.com/gravitl/netmaker/models"
)

//CreateUser creates a new user from
func CreateUser(c *gin.Context) {
	var user models.User
	fmt.Println("creating new user")
	user.UserName = c.PostForm("user")
	user.Password = c.PostForm("pass")
	if c.PostForm("admin") == "true" {
		user.IsAdmin = true
	} else {
		user.IsAdmin = false
	}
	user.Networks, _ = c.GetPostFormArray("network[]")
	fmt.Println("networks: ", user.Networks)
	_, err := controller.CreateUser(user)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, err, "Networks")
		return
	}
	ReturnSuccess(c, "Networks", "New user "+user.UserName+" created")
}

//DeleteUser delete a user
func DeleteUser(c *gin.Context) {
	user := c.PostForm("user")
	success, err := controller.DeleteUser(user)
	if !success {
		ReturnError(c, http.StatusBadRequest, err, "Networks")
		return
	}
	ReturnSuccess(c, "Networks", "user "+user+" deleted")
}

//EditUser displays form to update current user
func EditUser(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username").(string)
	user, err := controller.GetUser(username)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, err, "Networks")
		return
	}
	c.HTML(http.StatusOK, "EditUser", user)
}

//UpdateUser updates user from EditUser form
func UpdateUser(c *gin.Context) {
	var new models.User
	username := c.Param("user")
	user, err := controller.GetUserInternal(username)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, err, "Networks")
		return
	}
	new.UserName = c.PostForm("username")
	new.Password = c.PostForm("password")
	_, err = controller.UpdateUser(new, user)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, err, "Networks")
		return
	}
	ReturnSuccess(c, "Networks", "user has been updated")
}

//NewUser creates a new user; will fail if user with admin
//permissions already exists
func NewUser(c *gin.Context) {
	var user models.User
	user.UserName = c.PostForm("user")
	user.Password = c.PostForm("pass")
	user.IsAdmin = true
	hasAdmin, err := controller.HasAdmin()
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, err, "")
		return
	}
	if hasAdmin {
		ReturnError(c, http.StatusUnauthorized, errors.New("Admin Exists"), "")
		return
	}
	_, err = controller.CreateUser(user)
	if err != nil {
		ReturnError(c, http.StatusUnauthorized, err, "")
		return
	}
	location := url.URL{Path: "/"}
	c.Redirect(http.StatusFound, location.RequestURI())
}
