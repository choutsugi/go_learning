package main

import (
	"fmt"
	"math"
)

// 基本数据类型
func main() {
	// 十进制打印为二进制
	n := 10
	fmt.Printf("%b\n", n) // 二进制：1010
	fmt.Printf("%d\n", n) // 十进制：10

	// 八进制
	m := 075
	fmt.Printf("%o\n", m) // 八进制：075
	fmt.Printf("%d\n", m) // 十进制：61

	// 十六进制
	f := 0xff
	fmt.Printf("%x\n", f) // 十六进制：ff
	fmt.Printf("%X\n", f) // 十六进制：FF
	fmt.Printf("%d\n", f) // 十进制：255

	// uint8
	var age uint8 = 255 // 范围：0~255
	fmt.Println(age)

	// 浮点数
	fmt.Printf("%f\n", math.Pi)   // 3.141593
	fmt.Printf("%.2f\n", math.Pi) // 3.14

	fmt.Println(math.MaxFloat32) // 3.4028234663852886e+38
	fmt.Println(math.MaxFloat64) // 1.7976931348623157e+308

	// 复数
	var com1 complex64 = 1 + 2i  // (1+2i)
	var com2 complex128 = 3 + 4i // (3+4i)
	fmt.Println(com1)
	fmt.Println(com2)

	// 布尔值
	var a bool
	fmt.Println(a) // false
	a = true
	fmt.Println(a) // true

	// 字符串
	s1 := "hello shinrin"
	s2 := "halo ohiua"
	fmt.Println(s1) // hello shinrin
	fmt.Println(s2) // halo ohiua

	// 字符串转义符
	fmt.Println("str := \"c:\\users\\code\\golang.exe\"") // str := "c:\user\code\golang.exe"

	// 多行字符串
	s3 := `第一行
第二行
第三行
`
	fmt.Println(s3)
}
