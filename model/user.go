package model

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
}
