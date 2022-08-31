package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   uint64
	Name string
}

func main() {
	users := []User{
		{
			ID:   123,
			Name: "linonon",
		},
		{
			ID:   456,
			Name: "ononlin",
		},
	}

	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users) // http正常，則輸出
	})
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "The user id is %s", id)
	})
	// r.POST("/users")  創建
	// r.DELETE("/users")  刪除
	// r.PUT("/users/123")  更新
	// r.PATCH("/users/123")  更新部分
	r.Run(":8080")
}
