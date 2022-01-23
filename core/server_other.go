//go:build !windows
// +build !windows

package core

import (
	"github.com/fvbock/endless"
	"picture/global"
)

func CreatServer(address string) global.GinServer {
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
		fmt.Printf("server err: %v", err)
	}
}
