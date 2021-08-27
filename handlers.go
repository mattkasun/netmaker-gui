package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisplayLanding(c *gin.Context) {
	var Data PageData
	Data.Init("Networks")
	c.HTML(http.StatusOK, "layout", Data)
}
