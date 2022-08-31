package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeCountMW(c *gin.Context) {
	start := time.Now()

	c.Next()
	end := time.Since(start)
	fmt.Println("Time Pass:", end)
}

func main() {
	r := gin.Default()

	r.Use(TimeCountMW)

	r.GET("/time3", SecondThree)
	r.GET("/time5", SecondFive)

	r.Run(":9900")
}

func SecondThree(c *gin.Context) {
	time.Sleep(time.Second * 3)
	c.String(http.StatusOK, "Time pass 3 sec")
}

func SecondFive(c *gin.Context) {
	time.Sleep(time.Second * 5)
	c.String(http.StatusOK, "Time pass 5 sec")
}
