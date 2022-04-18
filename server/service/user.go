package service

import (
	"fmt"
	"whimsy/global"
	"whimsy/model"
	"whimsy/utils"
)

type UserService struct{}

// Register 用户注册
func (userService *UserService) Register(userModel model.User) (err error, userInter model.User) {

	return nil, model.User{}
}

// Login 用户登录
func (userService *UserService) Login(userModel *model.User) (err error, userInter *model.User) {
	if global.PIC_DB == nil {
		return fmt.Errorf("数据库未连接"), nil
	}
	//var user model.User
	userModel.Password, _ = utils.PasswordHash(userModel.Password)
	return nil, nil
}
