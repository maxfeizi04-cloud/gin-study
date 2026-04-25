package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maxfeizi04-cloud/gin-study/gin-web/res"
)

//type Response struct {
//	Code int         `json:"code"`
//	Msg  string      `json:"msg"`
//	Data interface{} `json:"data"`
//}
//
//func Index(c *gin.Context) {
//	c.JSON(http.StatusOK, Response{
//		Code: http.StatusOK,
//		Msg:  "Hello World!",
//		Data: nil,
//	})
//}

func main() {
	// gin 的初始化
	gin.SetMode("release")
	r := gin.Default()

	// 挂载路由
	r.GET("/index", func(c *gin.Context) {
		res.OkWithMsg(c, "肥子重复")
	})

	r.GET("/login", func(c *gin.Context) {
		res.OkWithData(c, map[string]any{
			"username": "admin",
			"password": "123456",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		res.FailWithData(c, 1001)
	})

	r.POST("/logout", func(c *gin.Context) {
		res.FailWithMsg(c, "登录错误")
	})

	// 绑定端口,启动服务
	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("服务启动失败:%s\n", err)
	}

}
