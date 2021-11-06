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
	fmt.Println(auth)
	c.Next()
}

func Files(c *gin.Context) {
	fmt.Println("display file upload form")
	c.HTML(http.StatusOK, "Files", nil)
}

func FileUpload(c *gin.Context) {
	fmt.Println("fileupload")
	filename := c.Param("file")
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
