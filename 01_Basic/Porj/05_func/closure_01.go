package main

import "fmt"

func Test() int {
	var x int
	x++
	return x
}

func main() {
	fmt.Println(Test()) // 1
	fmt.Println(Test()) // 1
	fmt.Println(Test()) // 1
}
