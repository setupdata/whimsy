package main

import (
	"whimsy/core"
	"whimsy/global"
)

func init() {
	// 初始化viper Viper是适用于Go应用程序的完整配置解决方案
	global.PIC_VIPER = core.InitViper()
	// 初始化日志库 logrus
	global.PIC_LOG = core.InitLogrus()
	// 初始化数据库 gorm
	global.PIC_DB = core.InitGorm()
	// 初始化路由
	global.PIC_ROUTER = core.InitRouter()
}

func main() {
	//启动服务
	core.InitServer()
	// 关闭前处理
	defer core.CloseServer()
}
