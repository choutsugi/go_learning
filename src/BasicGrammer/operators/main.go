package main

import "fmt"

// 运算符
func main() {
	// 算数运算符
	a := 20
	b := 10
	fmt.Println(a + b) // 30
	fmt.Println(a - b) // 10
	fmt.Println(a * b) // 200
	fmt.Println(a / b) // 2
	fmt.Println(a % b) // 0
	a++
	b--
	fmt.Println(a, b) // 21 9

	// 关系运算符
	fmt.Println(10 > 2)  // true
	fmt.Println(10 != 2) // true
	fmt.Println(4 <= 5)  // true

	// 逻辑运算符
	fmt.Println(10 > 5 && (5-4) == 1) // true
	fmt.Println(!(10 > 5))            // false
	fmt.Println(1 > 5 || (1+2) == 3)  // true

	// 位运算符
	c := 1              // 0001
	d := 5              // 0101
	fmt.Println(c & d)  // 0001
	fmt.Println(c | d)  // 0101
	fmt.Println(c ^ d)  // 0100
	fmt.Println(c << 2) // 0100
	fmt.Println(d >> 1) // 0010

	// 赋值运算符
	var e int
	e = 5
	e += 5 // e = e + 5
	fmt.Println(e)

}
