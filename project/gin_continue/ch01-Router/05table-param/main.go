package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/form", func(ctx *gin.Context) {
		types := ctx.DefaultPostForm("type", "post")
		username := ctx.PostForm("username")
		passord := ctx.PostForm("password")

		ctx.String(http.StatusOK, fmt.Sprintf("username:%s, password:%s, type: %s", username, passord, types))
	})

	r.Run(":9900")
}
