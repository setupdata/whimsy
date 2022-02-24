package core

import (
	"whimsy/global"
	"whimsy/initialize"
)

func InitServer() {
	global.PIC_LOG.Debug("启动服务器")
	//address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	ip := global.PIC_CONFIG.System.Ip
	address := global.PIC_CONFIG.System.Addr
	global.PIC_LOG.Info("监听地址: ", ip+":"+address)
	// 创建服务
	global.PIC_SERVER = initialize.CreatServer(ip + ":" + address)
	// 启动服务
	initialize.Listen()
}

func CloseServer() {
	global.PIC_LOG.Debug("清理数据库连接")
	if global.PIC_DB != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.PIC_DB.DB()
		_ = db.Close()
	}
}
