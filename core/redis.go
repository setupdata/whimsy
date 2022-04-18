package core

import (
	"context"
	"github.com/go-redis/redis/v8"
	"whimsy/global"
)

func InitRedis() *redis.Client {
	redisConfig := global.PIC_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.PIC_LOG.Error("redis连接测试错误: ", err)
		return nil
	} else {
		global.PIC_LOG.Info("redis连接测试成功")
		return client
	}
}
