package main

import (
	"ginstart/ch06/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/moreJSON", moreJSON)
	router.GET("/someProtoBuf", returnProto)
	router.Run(":8080")
}
func returnProto(c *gin.Context) {
	course := []string{"python", "go", "protobuf"}
	user := &proto.Teacher{
		Name:   "bobby",
		Course: course,
	}
	c.ProtoBuf(http.StatusOK, user)
}

func moreJSON(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"` // 配置轉換成json， Name -> user
		Message string
		Number  int
	}
	msg.Name = "bobby"
	msg.Message = "這是一個測試JSON"
	msg.Number = 20

	c.JSON(http.StatusOK, msg)
}
