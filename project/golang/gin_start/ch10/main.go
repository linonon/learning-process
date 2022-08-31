package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	// 當我們關閉程序的時候應該做的後續處理
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
	go func() {
		router.Run(":8080")
	}()
	// 如果想要接受信號
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 接收到退出信號
	fmt.Println("\nClosing server...")
	fmt.Println("\nbyebye")
}
