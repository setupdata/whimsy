//go:build !windows
// +build !windows

package initialize

import (
	"github.com/fvbock/endless"
	"time"
	"whimsy/global"
)

func CreatServer(address string) global.GinServer {
	global.PIC_LOG.Debug("创建http服务")
	s := endless.NewServer(address, global.PIC_ROUTER)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func Listen() {
	// 优雅地重启或停止
	err := global.PIC_SERVER.ListenAndServe()
	if err != nil {
		global.PIC_LOG.Fatal("http服务错误: %v", err)
	}
	global.PIC_LOG.Info("http服务结束")
}
