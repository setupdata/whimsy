package model

import (
	"whimsy/global"
)

type Vcode struct {
	global.BaseModel
	Username    string `json:"username" gorm:"comment:用户登录名"`
	Password    string `json:"password" gorm:"comment:用户登录密码"`
	NickName    string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	Avatar      string `json:"avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
	AuthorityId string `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
}
