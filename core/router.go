package core

import (
	"github.com/gin-gonic/gin"
	"whimsy/global"
	"whimsy/server/middleware"
	"whimsy/server/router"
)

//初始化Gin框架

func InitGin() *gin.Engine {
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

	global.PIC_LOG.Info("初始化中间件")
	//处理跨域请求
	r.Use(middleware.Cors()) // 直接放行全部跨域请求
	global.PIC_LOG.Debug("中间件-跨域")

	// 获取路由组实例
	sysRouter := router.RouterGroupApp
	// 创建路由分组
	PublicGroup := r.Group("")
	global.PIC_LOG.Debug("创建路由分组：PublicGroup")
	{ // 公有路由分组，无需验证
		sysRouter.UserRouter.InitPublicUserRouter(PublicGroup) // 注册共有用户路由
	}
	// 创建私有路由分组
	PrivateGroup := r.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	global.PIC_LOG.Debug("创建路由分组：PrivateGroup-中间件-JWTAuth")
	{ // 私有路由，需要jwt验证
		sysRouter.UserRouter.InitPrivateUserRouter(PrivateGroup) // 注册私有用户路由
	}

	//默认路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "Hello World",
		})
	})
	global.PIC_LOG.Info("路由初始化完成")
	return r
}
