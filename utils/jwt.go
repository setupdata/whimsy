package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"time"
	"whimsy/global"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("令牌已过期")
	TokenNotValidYet = errors.New("令牌尚未签发")
	TokenMalformed   = errors.New("这不是一个令牌")
	TokenInvalid     = errors.New("无法处理此令牌")
)

// BaseClaims 基础jwt荷载结构，自定义字段
type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	UserName    string
	AuthorityId string
}

// CustomClaims 用户jwt荷载结构
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

// NewJWT JWT构造函数
func NewJWT() *JWT {
	return &JWT{
		[]byte(global.PIC_CONFIG.JWT.SigningKey),
	}
}

// CreateClaims 创建JWT荷载结构
func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: global.PIC_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.PIC_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    global.PIC_CONFIG.JWT.Issuer,                          // 签名的发行者
		},
	}
	return claims
}

// CreateToken 创建并签发token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 缓冲期中的token以旧换新，回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
	v, err, _ := global.PIC_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// 令牌过期
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
