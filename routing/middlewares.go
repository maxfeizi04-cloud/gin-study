package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	fmt.Println("Home")
	fmt.Println(c.Get("Gm1"))
	fmt.Println(c.Get("Gm2"))
	//fmt.Println(c.Get("username"))
	userName, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username not found"})
	}
	fmt.Println(userName)

	c.String(http.StatusOK, "Hello World")
}

func M1(c *gin.Context) {
	fmt.Println("M1")
	c.Next()
	//c.Abort()
	fmt.Println("M1 End")
}

func M2(c *gin.Context) {
	fmt.Println("M2")
	c.Next()
	fmt.Println("M2 End")
}

func Gm1(c *gin.Context) {
	fmt.Println("Gm1 before")
	c.Set("Gm1", "John")
	var u UserInfo
	u.Name = "肥子"
	c.Set("username", u.Name)
	c.Next()
	fmt.Println("Gm1 after")
}

func Gm2(c *gin.Context) {
	fmt.Println("Gm2 before")
	c.Set("Gm2", "Tim")
	c.Next()
	fmt.Println(c.Get("Gm2"))
	fmt.Println("Gm2 after")
}

func main() {
	r := gin.Default()

	g := r.Group("/api")
	g.Use(Gm1, Gm2)
	//g.Use(Gm1)
	//g.Use(Gm2)

	g.GET("/users", Home)

	r.Run(":8080")
}

type UserInfo struct {
	Name string
}
