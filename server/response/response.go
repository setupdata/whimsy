package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 请求中的返回

var Code = map[string]int{
	"ok":        0,
	"reqErr":    400000, // 请求参数错误
	"tokenErr":  500000, // token异常
	"tokenErr1": 500001, // 未登录或非法访问
	"tokenErr2": 500002, // 令牌已过期
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// Ok 成功
func Ok(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(code int, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}

func OkWithData(code int, data interface{}, c *gin.Context) {
	Result(code, data, "操作成功", c)
}

func OkWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}

// Fail 失败
func Fail(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, "操作失败", c)
}

// FailWithMessage 失败-信息
func FailWithMessage(code int, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}

// FailWithDetailed 失败-详细
func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}
