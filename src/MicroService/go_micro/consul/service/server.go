// 创建服务
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello Web.")
	_, err := fmt.Fprintf(w, "Hello Web.")
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("health check.")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler)
	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
}
