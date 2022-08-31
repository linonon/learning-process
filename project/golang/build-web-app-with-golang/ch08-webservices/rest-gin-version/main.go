package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	router.GET("/user/:uid", getuser)
	router.POST("/adduser/:uid", adduser)
	router.DELETE("/deluser/:uid", deluser)
	router.PUT("/moduser/:uid", modifyuser)

	go func() {
		router.Run(":9900")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("")
	fmt.Println("Closing server")
	fmt.Println("Bye bye...")
}

func Index(c *gin.Context) {
	fmt.Fprint(c.Writer, "Welcome!\n")
}

func Hello(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Hello, %s\n", c.Param("name"))
}

func getuser(c *gin.Context) {
	uid := c.Param("uid")
	fmt.Fprintf(c.Writer, "You are get user %s", uid)
}

func modifyuser(c *gin.Context) {
	uid := c.Param("uid")
	fmt.Fprintf(c.Writer, "You are modify user %s", uid)
}

func deluser(c *gin.Context) {
	uid := c.Param("uid")
	fmt.Fprintf(c.Writer, "You are delete user %s", uid)
}

func adduser(c *gin.Context) {
	uid := c.Param("uid")
	fmt.Fprintf(c.Writer, "You are add user %s", uid)
}
