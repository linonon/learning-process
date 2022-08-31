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
	// URI version
	// r.GET("/:user/:password", func(c *gin.Context) {
	r.POST("/loginForm", func(c *gin.Context) {
		var form Login

		// URI version
		// if err := c.ShouldBindUri(&login); err != nil {
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "root" || form.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotModified})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": 200})
	})

	r.Run(":9900")
}
