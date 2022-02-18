package core

import (
	"gorm.io/gorm"
	"picture/global"
	"picture/initialize"
)

func InitGorm() *gorm.DB {
	switch global.PIC_CONFIG.System.DbType {
	case "mysql":
		return initialize.InitGormMysql()
	//case "pgsql":
	//return GormPgSql()
	default:
		return initialize.InitGormMysql()
	}
}
