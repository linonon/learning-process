package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusInternalServerError, "Upload file Error.")
		}

		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})

	r.Run(":9900")
}
