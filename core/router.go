package core

import (
	"github.com/gin-gonic/gin"
	"picture/global"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	//默认路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "Hello World",
		})
	})
	global.PIC_LOG.Error("路由初始化完成")
	return r
}
