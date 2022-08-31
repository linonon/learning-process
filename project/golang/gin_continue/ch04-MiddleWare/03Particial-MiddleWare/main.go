package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// All request will do MiddleWare() first
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("MiddleWare is running")
		c.Set("request", "Middle Ware")

		c.Next() // Like defer

		status := c.Writer.Status()
		fmt.Println("MiddleWare has finished", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	r := gin.Default()
	// r.Use(MiddleWare())

	// Particial MiddleWare
	r.GET("/ce", MiddleWare(), func(c *gin.Context) {
		req, _ := c.Get("request")
		fmt.Println("request:", req)

		c.JSON(http.StatusOK, gin.H{"request": req})
	})

	r.Run(":9900")
}
