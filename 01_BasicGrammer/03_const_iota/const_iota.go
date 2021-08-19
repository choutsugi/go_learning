/*
	常量。
*/
package main

import "fmt"

const (
	//iota首行为0，此后每行+1
	TEEMO = iota * 10 // TEEMO =  0, iota = 0
	YASUO             // YASUO = 10, iota = 1
	LUX               // LUX   = 20, iota = 2
	ZOE               // ZOE   = 30, iota = 3
)

const (
	a, b = iota + 1, iota + 2 // iota = 0, a = 1, b = 2
	c, d                      // iota = 1, a = 2, b = 3
	e, f                      // iota = 2, a = 3, b = 4
	g, h = iota * 2, iota * 3 // iota = 3, a = 6, b = 9
	i, j                      // iota = 4, a = 8, b = 12
)

func main() {
	// 常量（只读属性）
	const length int = 10
	fmt.Println("length = ", length)
	// length = 100 // 非法：常量无法修改

	fmt.Println("TEEMO = ", TEEMO)
	fmt.Println("YASUO = ", YASUO)
	fmt.Println("LUX = ", LUX)
	fmt.Println("ZOE = ", ZOE)
}
