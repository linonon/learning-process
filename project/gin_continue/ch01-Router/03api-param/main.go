package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/*aciton", func(c *gin.Context) {
		name := c.Param("name")
		aciton := c.Param("aciton")

		aciton = strings.Trim(aciton, "/")
		c.String(http.StatusOK, name+" is "+aciton)
	})

	r.Run(":9900")
	// http://127.0.0.1:9900/user/linonon/playing
}
