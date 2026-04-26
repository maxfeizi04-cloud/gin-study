package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// Static("别名", "实际文件目录路径")
	r.Static("st", "static")

	// StaticFile("别名", "具体文件的单个路径")
	r.StaticFile("abc", "static/abc.txt")

	r.Run(":8080")
}
