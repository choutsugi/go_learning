/*
	函数。
*/
package main

import "fmt"

// 匿名返回值
func fun1(a string, b int) int {
	fmt.Println("\n---func1---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 2333
	return c
}

// 匿名返回值
func fun2(a string, b int) (int, int) {
	fmt.Println("\n---func2---")

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 2, 333
}

// 具名返回值
func fun3(a string, b int) (ret1 int, ret2 int) {
	fmt.Println("\n---func3---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	ret1 = 2
	ret2 = 333
	return
}

func main() {
	// fun1
	c := fun1("BOOM", 7)
	fmt.Println("c = ", c)

	// fun2
	ret1, ret2 := fun2("BOOM", 7)
	fmt.Println("ret1 = ", ret1, ", ret2 = ", ret2)

	// fun3
	ret1, ret2 = fun2("BOOM", 7)
	fmt.Println("ret1 = ", ret1, ", ret2 = ", ret2)
}
