package router

import (
	"github.com/gin-gonic/gin"
	v1 "whimsy/server/api/v1"
)

type VCodeRouter struct{}

func (v *VCodeRouter) InitPublicVCodeRouter(Router *gin.RouterGroup) {
	// 获取业务逻辑
	vCodeApi := v1.ApiGroupApp.VCodeApi
	{
		Router.GET("getVCode", vCodeApi.GetVCode) // 获取验证码
	}
}
