package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"username" json:"user" uri:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"user" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "root" || json.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	r.Run(":9900")
}
