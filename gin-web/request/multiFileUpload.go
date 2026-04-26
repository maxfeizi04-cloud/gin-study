package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			log.Fatal(err)
			return
		}
		for _, header := range form.File {
			for _, file := range header {
				err = c.SaveUploadedFile(file, "upload/"+file.Filename)
				if err != nil {
					log.Fatal(err)
					return
				}
			}
		}
	})

	r.Run(":8080")
}
