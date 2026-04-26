package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		name := c.Query("name")
		query := c.DefaultQuery("age", "25")
		values := c.QueryArray("key")

		fmt.Println(name, query, values)
	})

	r.Run(":8080")
}
