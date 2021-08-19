package main

import (
	"fmt"
)

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	// interface{} 的类型断言机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type!")
	} else {
		fmt.Println("arg is string type! value = ", value)
		fmt.Printf("value type is %T\n", value)
	}

}

type Book struct {
	author string
}

func main() {
	book := Book{"Golang"}
	myFunc(book.author)
	myFunc(100)
}
