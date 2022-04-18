package service

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
	"whimsy/global"
	"whimsy/model"
	"whimsy/utils"
)

type PublicService struct{}

var ctx = context.Background()

// CreatVCode 生成验证码
func (p PublicService) CreatVCode() model.VCode {
	//生成验证码id
	newUuid := uuid.NewV4().String()
	//生成验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	numCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	//加入redis
	_, err := global.PIC_REDIS.SetNX(ctx, newUuid, numCode, 5*time.Minute).Result()
	if err != nil {
		global.PIC_LOG.Error("创建验证码时redis错误: ", err)
	}
	vCode := model.VCode{Id: newUuid, Num: numCode}
	return vCode
}

// SendVCode 发送验证码
func (p PublicService) SendVCode(to string, code string) {
	// 发送邮件
	utils.SendEmailWithTLS(to, code)
	return
}
