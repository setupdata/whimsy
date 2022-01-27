package global

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"picture/config"
)

//windows和linux服务抽象
type GinServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

// 组件实例
var (
	// PIC_CONFIG PIC_DB *gorm.DB
	//PIC_config
	PIC_CONFIG config.Server  // 配置文件
	PIC_LOG    *logrus.Logger // logrus日志实例
	PIC_VIPER  *viper.Viper   // viper实例
	PIC_ROUTER *gin.Engine    // 路由实例
	PIC_SERVER GinServer      // 服务实例
)

// 默认参数
var ()
