package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(ctx *gin.Context) {
		name := ctx.DefaultQuery("name", "linonon")
		ctx.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})

	r.Run(":9900")
}
