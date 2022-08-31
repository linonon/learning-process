package main

import (
	"ginc/ch03-Rendering/01EachTypeRen/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// TODO: 1.Json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "someJSON", "status": 200})
	})

	// TODO: 2. Struct
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}

		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})

	// TODO: 3.XML
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "someXML", "status": 200})
	})

	// TODO: 4.YAML
	r.GET("/someYAML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "someYAML", "status": 200})
	})

	// TODO: 5.Protobuf
	r.GET("/someProtobuf", func(c *gin.Context) {
		resp := []int64{int64(1), int64(2), int64(3)}

		label := "label"

		data := &pb.Test{
			Label: label,
			Resp:  resp,
		}

		c.ProtoBuf(http.StatusOK, data)
	})
}
