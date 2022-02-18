package main

import (
	"picture/core"
	"picture/global"
)

func init() {
	// 初始化viper Viper是适用于Go应用程序的完整配置解决方案
	global.PIC_VIPER = core.InitViper()
	// 初始化日志库 logrus
	global.PIC_LOG = core.InitLogrus()
	// 初始化数据库 gorm
	global.PIC_DB = core.InitGorm()
	//if global.PIC_DB == nil {
	//	panic(fmt.Errorf("数据库错误"))
	//}
	// 初始化路由
	global.PIC_ROUTER = core.InitRouter()
}

func main() {
	//启动服务
	core.InitServer()
	defer closeServer()
}

func closeServer() {
	if global.PIC_DB != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.PIC_DB.DB()
		_ = db.Close()
	}
}
