package initialize

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"picture/global"
	"time"
)

// MyWriter 定义自己的Writer
type MyWriter struct {
	myLog *log.Logger
}

// Printf 实现gorm/logger.Writer接口
func (m *MyWriter) Printf(format string, v ...interface{}) {
	logStr := fmt.Sprintf(format, v...)
	//利用logrus记录日志
	m.myLog.Debug(logStr)
}

// InitGormMysql 初始化mysql数据库
func InitGormMysql() *gorm.DB {
	m := global.PIC_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}

	gromLogger := logger.New(
		//设置Logger
		&MyWriter{myLog: global.PIC_LOG},
		logger.Config{
			//慢SQL阈值
			SlowThreshold: time.Millisecond,
			//设置日志级别，只有Warn以上才会打印sql
			LogLevel: logger.Warn,
		},
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度 utf8mb4
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: gromLogger,
	})

	if err != nil {
		global.PIC_LOG.Error(err)
		return nil
	}
	global.PIC_LOG.Debug(db)

	return db
}
