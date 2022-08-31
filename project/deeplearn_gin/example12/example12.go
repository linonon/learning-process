package main

import (
	"net/http"

	moudle "../../pb/pb"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/protobuf", func(c *gin.Context) {
		data := &moudle.User{
			Name: "linonon",
			Age:  20,
		}
		c.ProtoBuf(http.StatusOK, data)
	})
}
