package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Username string    `json:"userName" gorm:"comment:用户登录名"`
	Password string    `json:"-"  gorm:"comment:用户登录密码"`
	NickName string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	Avatar   string    `json:"avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
}
