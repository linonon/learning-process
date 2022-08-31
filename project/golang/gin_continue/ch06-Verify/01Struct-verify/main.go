package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	r := gin.Default()
	r.GET("/5lmh", func(c *gin.Context) {
		var person Person
		// gin 可以直接進行結構體認證
		if err := c.ShouldBind(&person); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprint(err))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("%#v", person))
	})

	r.Run(":9900")
}
