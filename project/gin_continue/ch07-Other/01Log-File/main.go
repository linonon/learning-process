package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()

	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(":9900")
}
