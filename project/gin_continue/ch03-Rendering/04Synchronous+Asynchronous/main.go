package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/long_async", func(c *gin.Context) {

		// Must use copied context.
		copyContext := c.Copy()

		go func() {
			time.Sleep(time.Second * 3)
			log.Println("Asynchronous request: " + copyContext.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(time.Second * 3)
		log.Println("Sync request: " + c.Request.URL.Path)
	})

	r.Run(":9900")
}
