package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("release")
	r := gin.Default()

	r.GET("/index", func(context *gin.Context) {

	})
	r.POST("/index")

	r.Run(":8080")
}
