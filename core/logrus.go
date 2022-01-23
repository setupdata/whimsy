package core

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLogrus() *logrus.Logger {
	log := logrus.New()
	// 日志输出位置
	log.SetOutput(os.Stdout)

	//设置日志级别
	log.SetLevel(logrus.InfoLevel)

	// 设置日志格式
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//log.SetFormatter(&logrus.JSONFormatter{})
	return log
}

//func newLfsHook(logLevel *string) logrus.Hook {
//	logName := "pic-log"
//	// WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
//	// WithRotationTime设置日志分割的时间,这里设置为一小时分割一次
//	// WithMaxAge和WithRotationCount二者只能设置一个,
//	// WithMaxAge设置文件清理前的最长保存时间,
//	// WithRotationCount设置文件清理前最多保存的个数.
//	writer, err := rotatelogs.New(
//		logName+".%Y%m%d%H",
//		rotatelogs.WithLinkName(logName),
//		rotatelogs.WithRotationTime(time.Hour*24),
//		rotatelogs.WithMaxAge(time.Hour*24*30),
//		//maxRemainCnt uint
//		//rotatelogs.WithRotationCount(maxRemainCnt),
//	)
//
//	if err != nil {
//		//log.Errorf("config local file system for logger error: %v", err)
//	}
//
//	level, ok := logLevels[*logLevel]
//
//	if ok {
//		log.SetLevel(level)
//	} else {
//		log.SetLevel(log.WarnLevel)
//	}
//
//	lfsHook := lfshook.NewHook(lfshook.WriterMap{
//		log.DebugLevel: writer,
//		log.InfoLevel:  writer,
//		log.WarnLevel:  writer,
//		log.ErrorLevel: writer,
//		log.FatalLevel: writer,
//		log.PanicLevel: writer,
//	}, &log.TextFormatter{DisableColors: true})
//
//	return lfsHook
//}

//func init() {
//	path := "/Users/opensource/test/go.log"
//	/* 日志轮转相关函数
//	`WithLinkName` 为最新的日志建立软连接
//	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
//	WithMaxAge 和 WithRotationCount二者只能设置一个
//	 `WithMaxAge` 设置文件清理前的最长保存时间
//	 `WithRotationCount` 设置文件清理前最多保存的个数
//	*/
//	// 下面配置日志每隔 1 分钟轮转一个新文件，保留最近 3 分钟的日志文件，多余的自动清理掉。
//	writer, _ := rotatelogs.New(
//		path+".%Y%m%d%H%M",
//		rotatelogs.WithLinkName(path),
//		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
//		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
//	)
//	log.SetOutput(writer)
//	//log.SetFormatter(&log.JSONFormatter{})
//}
