package main

import "fmt"

func main() {

	// 声明变量
	var f1 float32
	f1 = 3.14
	fmt.Println("f1 = ", f1)

	// 自动推导类型
	f2 := 3.14
	fmt.Println("f2 = ", f2)

	// 自动推导时通过.推到为浮点型
	f3 := 3.
	fmt.Printf("f3 = %f", f3)
}
