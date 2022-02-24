package core

import (
	"bytes"
	"fmt"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"time"
	"whimsy/global"
)

// MyFormatter 自定义输出格式
type MyFormatter struct{}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//global.PIC_LOG.Debug(entry, "\n")
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
		return b.Bytes(), fmt.Errorf("输入为空")
	}
	// 时间格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	var newLog string
	//HasCaller()为true才会有调用信息
	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[%s] [%s] %s [%s:%d %s] \n",
			timestamp, entry.Level, entry.Message, fName, entry.Caller.Line, entry.Caller.Function)
	} else {
		newLog = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}

// 日志文件分割
func creatWriter() *rotateLogs.RotateLogs {
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
	writer, err := rotateLogs.New(
		logPath+logName+"-%Y%m%d%H%M."+logType,
		rotateLogs.WithLinkName(baseLogPath),                                                       // 生成软链，指向最新日志文件，软链接Symlink在windows下需要权限
		rotateLogs.WithMaxAge(time.Minute*time.Duration(global.PIC_CONFIG.Log.MaxAge)),             // 文件最大保存时间，最小分钟为单位
		rotateLogs.WithRotationTime(time.Minute*time.Duration(global.PIC_CONFIG.Log.RotationTime)), // 日志切割时间间隔，最小为1分钟轮询，默认1分钟，低于1分钟就按1分钟来
	)

	if err != nil {
		panic(fmt.Errorf("配置本地日志服务错误: %v", errors.WithStack(err)))
	}

	return writer
}

func InitLogrus() *logrus.Logger {
	myLogger := logrus.New()

	//设置日志级别
	level := global.PIC_CONFIG.Log.Level
	if global.PIC_CONFIG.System.Env == "develop" {
		// 开发环境
		level = "Debug"
	}
	switch level {
	case "Debug":
		myLogger.SetLevel(logrus.DebugLevel)
	case "Info":
		myLogger.SetLevel(logrus.InfoLevel)
	case "Warn":
		myLogger.SetLevel(logrus.WarnLevel)
	case "Error":
		myLogger.SetLevel(logrus.ErrorLevel)
	case "Fatal":
		myLogger.SetLevel(logrus.FatalLevel)
	case "Panic":
		myLogger.SetLevel(logrus.PanicLevel)
	default:
		myLogger.SetLevel(logrus.InfoLevel)
	}

	if level != "Info" {
		// 等级不为Info时，输出文件名、行号及函数名
		myLogger.SetReportCaller(true)
	}

	// 设置日志格式
	switch format := global.PIC_CONFIG.Log.Format; format {
	case "text":
		myLogger.SetFormatter(&logrus.TextFormatter{
			ForceQuote:      true, // 输出加双引号
			TimestampFormat: "2006-01-02 15:04:05",
		})
	case "json":
		myLogger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	case "piclog":
		myLogger.SetFormatter(&MyFormatter{})
	default:
		myLogger.SetFormatter(&MyFormatter{})
	}

	// 日志输出方式
	myLogger.SetOutput(creatWriter())

	// 日志
	myLogger.Info("logrus初始化完成")
	return myLogger
}
