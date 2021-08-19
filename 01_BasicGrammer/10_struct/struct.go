package main

import "fmt"

// 通过type指定别名
type _int int

// 定义结构体
type Book struct {
	title string
	auth  string
}

// 值传递：副本
func changeBook1(book Book) {
	book.auth = "Jax"
}

// 引用传递：
func changeBook2(book *Book) {
	book.auth = "Master Yi"
}

func main() {
	var book Book
	book.title = "Golang"
	book.auth = "Master"

	fmt.Printf("%v\n", book)
	changeBook1(book)
	fmt.Printf("%v\n", book)
	changeBook2(&book)
	fmt.Printf("%v\n", book)
}
