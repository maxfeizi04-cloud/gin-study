package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 取得动态参数
	r.GET("/user/:id/:name", func(c *gin.Context) {
		userid := c.Param("id")
		username := c.Param("name")
		fmt.Println(userid, username)
	})

	r.Run(":8080")
}
