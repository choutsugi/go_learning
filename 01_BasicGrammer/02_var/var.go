/*
	变量。
*/

package main

import "fmt"

//声明全局变量
var global_A int = 100
var global_B = 200

// global_C:= 300 //非法，:=仅用于函数体内

func main() {
	//方法一：声明一个变量，默认值为0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	//方法二：声明一个变量，默认值为0
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	//方法三：初始化时省去数据类型，通过值自动匹配当前变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	//方法四（常用，仅用于函数体内）：省去var关键字，自动匹配
	d := 100
	fmt.Println("d = ", d)
	fmt.Printf("type of d = %T\n", d)

	dd := "abcd"
	fmt.Println("dd = ", dd)
	fmt.Printf("type of dd = %T\n", dd)

	//打印全局变量
	fmt.Println("global_A = ", global_A, ", global_B = ", global_B)
	// fmt.Println("global_C = ", global_C)

	//声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, ", yy = ", yy)
	var zz, ss = 300, "abcd"
	fmt.Println("zz = ", zz, ", ss = ", ss)

	//多行多变量声明
	var (
		mm int  = 400
		tt bool = true
	)
	fmt.Println("mm = ", mm, ", tt = ", tt)
}
