package main

import "fmt"

func main() {
	// 1.声明bool变量但不初始化，默认值为false
	var a bool
	fmt.Println("a = ", a)

	a = true
	fmt.Println("a = ", a)

	// 2.自动推导类型
	var b = false
	fmt.Println("b = ", b)

	c := true
	fmt.Println("c = ", c)

	// 格式化输出
	var d bool
	d = false
	fmt.Printf("d = %t", d)
}
