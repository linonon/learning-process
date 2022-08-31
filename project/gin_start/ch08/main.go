package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("example", "123456")
		c.Next()

		end := time.Since(t)
		fmt.Printf("耗時 %v\n", end)

		status := c.Writer.Status()
		fmt.Println("狀態：", status)
	}
}

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		for k, v := range c.Request.Header {
			fmt.Println("p rint:", k, v, token)
			if k == "X-Token" {
				token = v[0]
			} else {
				fmt.Println(k, v)
			}
		}
		fmt.Println("linonon?", token)
		if token != "linonon" {
			c.JSON(401, gin.H{
				"message": "fail",
			})
		}
		c.Next()
	}
}
func main() {
	router := gin.Default()
	router.Use(Token(), MyLogger())

	router.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	router.Run(":8080")

}
