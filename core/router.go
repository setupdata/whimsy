package core

import (
	"github.com/gin-gonic/gin"
	"picture/global"
)

func InitRouter() *gin.Engine {
	// 设置gin模式
	switch ginMod := global.PIC_CONFIG.System.GinMod; ginMod {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	//默认路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "Hello World",
		})
	})
	global.PIC_LOG.Info("路由初始化完成")
	return r
}
