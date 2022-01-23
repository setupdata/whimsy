package core

import "picture/global"

func InitServer() {
	//address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	address := ":8080"
	// 创建服务
	global.PIC_SERVER = CreatServer(address)
	// 启动服务
	Listen()
}
