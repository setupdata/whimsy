package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"whimsy/server/request"
	"whimsy/server/response"
)

type UserApi struct{}

func (u *UserApi) Register(c *gin.Context) {
	fmt.Print("用户注册")
	//l := request.RequestGroupApp.User
	response.OkWithMessage(response.Code["ok"], "一切正常", c)
}

// Login 用户登录
func (u *UserApi) Login(c *gin.Context) {
	fmt.Print("用户登录\n")
	//检查请求参数是否符合结构
	l := request.RequestGroupApp.User.Login
	if err := c.ShouldBindJSON(&l); err != nil {
		// 请求参数和结构体不匹配
		response.FailWithMessage(response.Code["reqErr"], "请求错误", c)
		return
	}

	//user := &model.User{
	//	Username: l.Username,
	//	Password: l.Password,
	//}

	response.OkWithMessage(response.Code["ok"], "一切正常", c)
}
