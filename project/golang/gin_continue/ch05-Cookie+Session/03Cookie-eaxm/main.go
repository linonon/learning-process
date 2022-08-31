package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("abc")
		if cookie == "123" {
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error:": "err"})
		c.Abort()
	}
}
func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		// Must use "localhost" instead of "127.0.0.1"
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		c.String(http.StatusOK, "Login success!")
	})

	r.GET("/home", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "home"})
	})

	r.Run(":9900")
}
