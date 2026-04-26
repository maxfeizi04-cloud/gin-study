package main

import (
	bytes2 "bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		bytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(bytes))
		// 读取之后 Body 就没数据了， 阅后即焚
		//value := c.PostForm("name")
		//fmt.Println(value)
		/*
			这段代码是把一个字节数组通过 bytes.NewReader 转成一个 Reader,
			再用 io.NopCloser 包装成 ReadCloser，
			用于满足一些接口（比如 http.Request.Body）对 io.ReadCloser 的要求。
			本质上是把内存数据伪装成一个可关闭的流。
		*/
		c.Request.Body = io.NopCloser(bytes2.NewReader(bytes))
		fmt.Println(c.Request.Header)
	})

	r.Run(":8080")
}
