package core

import (
	"gorm.io/gorm"
	"whimsy/global"
	"whimsy/initialize"
)

// 建立mysql数据库连接
func creatMysql() *gorm.DB {
	// 建立连接
	sqlDB := initialize.InitGormMysql()
	if sqlDB != nil {
		// 迁移数据库表
		initialize.InitTales(sqlDB)
	} else {
		global.PIC_LOG.Error("数据库错误: db对象为空")
	}
	return sqlDB
}

func InitGorm() *gorm.DB {
	global.PIC_LOG.Info("初始化数据库")
	switch global.PIC_CONFIG.System.DbType {
	case "mysql":
		return creatMysql()
	//case "pgsql":
	//return GormPgSql()
	default:
		return creatMysql()
	}
}
