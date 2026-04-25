package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.String())
	if r.Method == "GET" || r.Method == "POST" {
		byteData, _ := io.ReadAll(r.Body)
		fmt.Println("这是请求 Body 的信息: ", string(byteData))
	}

	fmt.Printf("这是请求 Header 的信息:%s\n", r.Header)

	k, err := w.Write([]byte("Hello World!"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(k)

}

func main() {
	http.HandleFunc("/index", Index)
	fmt.Printf("http server running http://localhost:8080/ \n")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
