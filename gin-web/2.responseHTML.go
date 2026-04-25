package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	// r.LoadHTMLFiles("templates/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]any{
			"title": "worladad",
		})
	})

	r.Run(":8080")
}
