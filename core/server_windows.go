//go:build windows
// +build windows

package core

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"picture/global"
	"time"
)

func CreatServer(address string) global.GinServer {
	global.PIC_LOG.Info("创建http服务")
	return &http.Server{
		Addr:           address,
		Handler:        global.PIC_ROUTER,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func Listen() {
	global.PIC_LOG.Info("启动http服务")
	// 优雅地重启或停止
	go func() {
		// service connections
		if err := global.PIC_SERVER.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := global.PIC_SERVER.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}