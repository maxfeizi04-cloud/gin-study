package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		type user struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}
		var u user
		err := c.ShouldBindQuery(&u)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(u)
	})

	r.GET("user/:id/:name", func(c *gin.Context) {
		type user struct {
			Name string `uri:"name"`
			Id   int    `uri:"id"`
		}
		var u user
		err := c.ShouldBindUri(&u)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(u)
	})

	// Form 参数
	r.POST("form", func(c *gin.Context) {
		type user struct {
			Name string `form:"name"`
			Id   int    `form:"id"`
		}
		var u user
		err := c.ShouldBind(&u)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(u)
	})

	// Json 参数
	r.POST("json", func(c *gin.Context) {
		type user struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		}
		var u user
		err := c.ShouldBindJSON(&u)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(u)
	})

	// Header 参数
	r.POST("header", func(c *gin.Context) {
		type user struct {
			Name      string `header:"name"`
			Id        int    `header:"id"`
			UserAgent string `header:"user-agent"`
		}
		var u user
		err := c.ShouldBindHeader(&u)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(u)
	})
	r.Run(":8080")
}
