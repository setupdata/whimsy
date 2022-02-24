//go:build windows
// +build windows

package initialize

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
	"whimsy/global"
)

func CreatServer(address string) global.GinServer {
	global.PIC_LOG.Debug("创建http服务")
	return &http.Server{
		Addr:           address,
		Handler:        global.PIC_ROUTER,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func Listen() {
	global.PIC_LOG.Info("http服务启动")
	// 优雅地重启或停止
	go func() {
		// service connections
		if err := global.PIC_SERVER.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// Fatal() 进程退出代码为 1
			global.PIC_LOG.Fatal("http服务错误: ", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	global.PIC_LOG.Debug("关闭http服务")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := global.PIC_SERVER.Shutdown(ctx); err != nil {
		global.PIC_LOG.Fatal("http服务关闭因为:", err)
	}
	global.PIC_LOG.Info("http服务结束")
}
