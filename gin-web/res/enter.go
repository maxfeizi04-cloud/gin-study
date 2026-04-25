package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var codeMap = map[int]string{
	1001: "权限错误",
	1002: "角色错误",
}

func response(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Ok(c *gin.Context, msg string, data interface{}) {
	response(c, 0, msg, data)
}

func OkWithData(c *gin.Context, data interface{}) {
	Ok(c, "成功", data)
}

func OkWithMsg(c *gin.Context, msg string) {
	Ok(c, msg, nil)
}

func Fail(c *gin.Context, code int, msg string, data interface{}) {
	response(c, code, msg, data)
}

func FailWithData(c *gin.Context, code int) {
	msg, ok := codeMap[code]
	if !ok {
		msg = "服务错误"
	}
	Fail(c, code, msg, nil)
}

func FailWithMsg(c *gin.Context, msg string) {

	Fail(c, 1001, msg, nil)
}
