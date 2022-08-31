package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	goodGroup := router.Group("/goods")
	{
		goodGroup.GET("", goodsList)
		//匹配URL，需要完全一致才能正常匹配
		// *action : 得到 '/' + 'action' + '...'
		// :action : 只會得到String(action)
		goodGroup.GET("/:id/:action", goodsDetail)
		goodGroup.POST("", createGoods)
	}
	router.Run(":8080")
}

func goodsList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "goodList",
	})
}

// 獲得 URL 當中的 id 值
func goodsDetail(c *gin.Context) {
	id := c.Param("id")
	action := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"action": action,
	})
}

func createGoods(c *gin.Context) {

}
