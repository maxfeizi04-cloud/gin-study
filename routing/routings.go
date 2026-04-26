package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserView(c *gin.Context) {
	path := c.Request.URL
	fmt.Println(c.Request.Method, path)
}

func main() {
	r := gin.Default()

	// r.GET() 获取资源
	// r.POST()  创建新资源
	// r.PUT() 更新整个资源
	// r.PATCH() 部分更新资源
	// r.DELETE() 删除资源
	// r.Any()
	api := r.Group("api")
	UserGroup(api)

	login := r.Group("api")
	loginGroup(login)

	r.Run(":8080")
}

func UserGroup(r *gin.RouterGroup) {
	r.GET("users", UserView)
	r.POST("users", UserView)
	r.DELETE("users", UserView)
	r.PUT("users", UserView)
	r.PATCH("users", UserView)
}

func loginGroup(r *gin.RouterGroup) {
	r.GET("login", UserView)
}
