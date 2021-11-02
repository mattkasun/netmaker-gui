package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	controller "github.com/gravitl/netmaker/controllers"
	"github.com/gravitl/netmaker/models"
)

func ProcessLogin(c *gin.Context) {
	fmt.Println("Processing Login")
	var AuthRequest models.UserAuthParams
	AuthRequest.UserName = c.PostForm("user")
	AuthRequest.Password = c.PostForm("pass")
	session := sessions.Default(c)
	//don't need the jwt
	_, err := controller.VerifyAuthRequest(AuthRequest)
	if err != nil {
		fmt.Println("error verifying AuthRequest: ", err)
		session.Set("message", err.Error())
		session.Set("loggedIn", false)
		c.HTML(http.StatusUnauthorized, "Login", gin.H{"message": err})
	} else {
		session.Set("loggedIn", true)
		//init message
		session.Set("message", "")
		session.Options(sessions.Options{MaxAge: 28800})
		user, err := controller.GetUser(AuthRequest.UserName)
		if err != nil {
			fmt.Println("err retrieving user: ", err)
		}
		session.Set("username", user.UserName)
		session.Set("isAdmin", user.IsAdmin)
		session.Set("networks", user.Networks)
		session.Save()
		location := url.URL{Path: "/"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
}

func DisplayLanding(c *gin.Context) {
	var data PageData
	var message string
	page := ""
	session := sessions.Default(c)
	if session.Get("page") != nil {
		page = session.Get("page").(string)
	}
	if session.Get("message") != nil {
		message = session.Get("message").(string)
	}
	fmt.Println("Initialializing PageData for page ", page, "with message ", message)
	if page != "" {
		data.Init(c, page, message)
	} else {
		data.Init(c, "Networks", message)
	}
	//clear message
	session.Set("message", "")
	session.Save()
	c.HTML(http.StatusOK, "layout", data)
}

func LogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("loggedIn", false)
	session.Set("message", "")
	session.Save()
	fmt.Println("User Logged Out", session.Get("loggedIn"))
	location := url.URL{Path: "/"}
	c.Redirect(http.StatusFound, location.RequestURI())
}


func ReturnSuccess(c *gin.Context, page, message string) {
	session := sessions.Default(c)
	session.Set("messsge", message)
	session.Set("page", page)
	session.Save()
	location := url.URL{Path: "/"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func ReturnError(c *gin.Context, status int, err error, page string) {
	var data PageData
	session := sessions.Default(c)
	session.Set("message", err.Error())
	session.Set("page", page)
	session.Save()
	if page != "" {
		data.Init(c, page, err.Error())
	} else {
		data.Init(c, "Networks", err.Error())
	}
	c.HTML(status, "layout", data)
}
