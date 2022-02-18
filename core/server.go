package core

import (
	"picture/global"
	"picture/initialize"
)

func InitServer() {
	//address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	ip := global.PIC_CONFIG.System.Ip
	address := global.PIC_CONFIG.System.Addr
	global.PIC_LOG.Info("监听地址: ", ip+":"+address)
	// 创建服务
	global.PIC_SERVER = initialize.CreatServer(ip + ":" + address)
	// 启动服务
	initialize.Listen()
}
