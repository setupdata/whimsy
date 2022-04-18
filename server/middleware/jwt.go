package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"whimsy/global"
	"whimsy/server/response"
	"whimsy/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("X-Token")
		if token == "" {
			// token为空需要登录
			response.FailWithMessage(response.Code["tokenErr1"], "未登录或非法访问", context)
			context.Abort()
			return
		}
		// 解析token包含的信息
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				// 授权已过期
				response.FailWithMessage(response.Code["tokenErr2"], "令牌已过期", context)
				context.Abort()
				return
			}
			// 其他错误
			response.FailWithMessage(response.Code["tokenErr"], err.Error(), context)
			context.Abort()
			return
		}
		// token到达缓冲时间
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			// 重新设置过期时间
			claims.ExpiresAt = time.Now().Unix() + global.PIC_CONFIG.JWT.ExpiresTime
			// 创建新令牌
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			//newClaims, _ := j.ParseToken(newToken)
			// 返回新令牌
			context.Header("New-Token", newToken)
			//context.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		context.Set("claims", claims)
		context.Next()
	}
}
