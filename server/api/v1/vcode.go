package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"whimsy/server/response"
)

type VCodeApi struct{}

func (v *VCodeApi) GetVCode(c *gin.Context) {
	fmt.Print("发送验证码\n")
	// 创建验证码
	vCode := publicService.CreatVCode()
	// 发送验证码
	publicService.SendVCode("1144639044@qq.com", vCode.Num)
	response.OkWithDetailed(response.Code["ok"], struct{ Id string }{Id: vCode.Id}, "一切正常", c)
}
