package main

import "fmt"

// myFunc
func myFunc(a, b int) (sum int) {
	sum = a + b
	return
}

// 函数类型（函数指针）
type FuncType func(a, b int) int

func main() {
	var result FuncType
	result = myFunc
	s := result(3, 5)
	fmt.Println(s)
}
