package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename=3.fileDownload.go")
		c.File("3.fileDownload.go")
	})

	r.Run(":8080")
}
