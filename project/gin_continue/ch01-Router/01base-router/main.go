package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.POST("/xxxpost", getting)
	r.PUT("/xxxput")

	r.Run(":9900")
}

func getting(c *gin.Context) {
	c.String(http.StatusOK, "hello poster")
}
