package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func FileAuth(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth != "Bearer secretkey" {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to view this endpoing"})
		return
	}
	c.Next()
}

func FileUpload(c *gin.Context) {
	filename := c.Param("file")
	//although this is using user provided input, it is not a security issue as routing will result in 404 error if path elements are included in filename
	file, err := os.Create("./files/" + filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	n, err := io.Copy(file, c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("upload %d bytes to file  %s", n, filename))
}
