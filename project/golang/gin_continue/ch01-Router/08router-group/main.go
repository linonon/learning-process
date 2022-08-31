package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "linonon")
	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "linonon submit")
	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
}
