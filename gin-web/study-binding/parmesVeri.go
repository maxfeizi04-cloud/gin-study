package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/json", func(c *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"required"`
			Age  int    `json:"age" binding:"min=1,max=150"`
		}
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user})
	})

	r.Run(":8080")
}
