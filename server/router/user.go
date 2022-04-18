package router

import (
	"github.com/gin-gonic/gin"
	v1 "whimsy/server/api/v1"
)

type UserRouter struct{}

func (s *UserRouter) InitPublicUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	// 获取业务逻辑
	userApi := v1.ApiGroupApp.UserApi
	{
		userRouter.POST("login", userApi.Login) // 用户登录
	}
}

func (s *UserRouter) InitPrivateUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userApi := v1.ApiGroupApp.UserApi
	{
		userRouter.POST("register", userApi.Register) // 用户注册
	}
}
