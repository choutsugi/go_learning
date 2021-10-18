package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	files, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:", err)
		return
	}
	// 渲染模板
	name := "stone"
	err = files.Execute(w, name)
	if err != nil {
		fmt.Println("Render template failed, err:", err)
		return
	}
}

func main() {

	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err:", err)
		return
	}

}
