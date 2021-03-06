package global

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"time"
	"whimsy/config"
)

// GinServer
// windows和linux服务抽象
// http.server和endlessServer接口
type GinServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

// BaseModel
// 基础数据库组成
type BaseModel struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

// 组件实例
var (
	// PIC_CONFIG PIC_DB *gorm.DB
	PIC_CONFIG              config.Server           // 配置文件
	PIC_LOG                 *logrus.Logger          // logrus日志实例
	PIC_VIPER               *viper.Viper            // viper实例
	PIC_DB                  *gorm.DB                // 数据库实例
	PIC_REDIS               *redis.Client           // redis实例
	PIC_GIN                 *gin.Engine             // Gin实例
	PIC_SERVER              GinServer               // 服务实例
	PIC_Concurrency_Control = &singleflight.Group{} // 并发控制 jwt以旧换新
)

// 默认参数
var ()
