package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")

		// 读取文件名
		log.Println(file.Filename)

		// 文件大小是字节 byte
		log.Println(file.Size)

		//open, _ := file.Open()
		//bytes, _ := io.ReadAll(open)

		//err := os.WriteFile("xxx.jpg", bytes, 0666)
		//if err != nil {
		//	log.Fatal(err)
		//}
		err := c.SaveUploadedFile(file, "upload/"+file.Filename)
		if err != nil {
			log.Fatal(err)
		}
	})

	r.Run(":8080")
}
