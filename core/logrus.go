package core

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"picture/global"
	"time"
)

func InitLogrus() *log.Logger {
	myLoger := log.New()

	//设置日志级别
	if global.PIC_CONFIG.System.Env == "develop" {
		// 开发环境
		myLoger.SetLevel(log.DebugLevel)
	} else {
		switch level := global.PIC_CONFIG.Log.Level; level {
		case "debug":
			myLoger.SetLevel(log.DebugLevel)
		case "info":
			myLoger.SetLevel(log.InfoLevel)
		case "warn":
			myLoger.SetLevel(log.WarnLevel)
		case "error":
			myLoger.SetLevel(log.ErrorLevel)
		default:
			myLoger.SetLevel(log.InfoLevel)
		}
	}

	// 设置日志格式
	switch format := global.PIC_CONFIG.Log.Format; format {
	case "text":
		myLoger.SetFormatter(&log.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	case "json":
		myLoger.SetFormatter(&log.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}

	/* 日志轮转相关函数
	 *WithLinkName` 为最新的日志建立软连接
	 *WithRotationTime` 设置日志分割的时间，隔多久分割一次
	 *WithMaxAge 和 WithRotationCount二者只能设置一个
	 *WithMaxAge` 设置文件清理前的最长保存时间
	 *WithRotationCount` 设置文件清理前最多保存的个数
	 */
	// 下面配置日志每隔 24小时 轮转一个新文件，保留最近 30天 的日志文件，多余的自动清理掉。
	logPath := global.PIC_CONFIG.Log.Path
	logName := global.PIC_CONFIG.Log.Name
	logType := global.PIC_CONFIG.Log.Type
	baseLogPath := logPath + logName + "." + logType
	writer, err := rotatelogs.New(
		logPath+logName+"-%Y%m%d%H%M."+logType,
		rotatelogs.WithLinkName(baseLogPath),                                                       // 生成软链，指向最新日志文件，软链接Symlink在windows下需要权限
		rotatelogs.WithMaxAge(time.Minute*time.Duration(global.PIC_CONFIG.Log.MaxAge)),             // 文件最大保存时间，最小分钟为单位
		rotatelogs.WithRotationTime(time.Minute*time.Duration(global.PIC_CONFIG.Log.RotationTime)), // 日志切割时间间隔，最小为1分钟轮询，默认60s，低于1分钟就按1分钟来
	)

	if err != nil {
		panic(fmt.Errorf("配置本地日志服务错误: %v", errors.WithStack(err)))
	}

	// 日志输出方式
	myLoger.SetOutput(writer)

	return myLoger
}
