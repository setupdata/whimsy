package initialize

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
	"whimsy/global"
	"whimsy/model"
)

// MyWriter 定义自己的Writer
type MyWriter struct {
	// logrus输出日志
	myLogrus *logrus.Logger
	// gorm默认logger输出日志，输出到控制台
	gormLogger logger.Writer
}

// Printf 实现gorm/logger.Writer接口
func (m *MyWriter) Printf(format string, v ...interface{}) {
	logStr := fmt.Sprintf(format, v...)
	if global.PIC_CONFIG.Mysql.LogLogrus {
		// 利用logrus记录日志
		m.myLogrus.Info(fmt.Sprintf(logStr))
	} else {
		// gorm默认logger输出日志，输出到控制台
		m.gormLogger.Printf(logStr)
	}
}

// 设置gorm config
func creatGormConfig() *gorm.Config {
	config := &gorm.Config{
		// 不跳过默认事务，保证数据稳定性
		SkipDefaultTransaction: false,
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "pic_", // 表名前缀，`User` 的表名应该是 `pic_users`
			SingularTable: false,  // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		// 迁移时禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true, // 不使用物理外键，使用逻辑外键，在代码中自动提现外键关系
	}
	gormLogger := logger.New(
		//设置Logger
		&MyWriter{
			myLogrus:   global.PIC_LOG,
			gormLogger: log.New(os.Stdout, "\r\n", log.LstdFlags),
		},
		logger.Config{
			SlowThreshold: time.Second,                        // 慢 SQL 阈值, 会在日志中标红, 并且标记 [SLOW SQL >= 1s]
			LogLevel:      logger.Warn,                        // Log level
			Colorful:      !global.PIC_CONFIG.Mysql.LogLogrus, // 由输出方向控制是否有颜色
		},
	)
	switch global.PIC_CONFIG.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = gormLogger.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = gormLogger.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = gormLogger.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = gormLogger.LogMode(logger.Info)
	default:
		config.Logger = gormLogger.LogMode(logger.Info)
	}
	return config
}

// InitGormMysql 初始化mysql数据库
func InitGormMysql() *gorm.DB {
	global.PIC_LOG.Debug("连接mysql数据库")
	m := global.PIC_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}

	// 连接数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度 utf8mb4
		SkipInitializeWithVersion: false,   // 根据当前 MySQL 版本自动配置
	}), creatGormConfig())

	if err != nil {
		global.PIC_LOG.Error(err)
		return nil
	} else {
		sqlDB, _ := db.DB()
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)

		return db
	}
}

// 迁移数据库
func InitTales(sqlDB *gorm.DB) {
	global.PIC_LOG.Debug("迁移数据库表")
	err := sqlDB.AutoMigrate(
		model.User{},
	)
	if err != nil {
		global.PIC_LOG.Error("迁移数据库表错误", err)
	}
}
