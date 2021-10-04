package main

import "fmt"

func fine() (int, string) {
	return 10, "HAKUNO"
}

// 变量
func main() {
	// 标准声明格式
	var name string // ""
	var age int     // 0
	var isOk bool   // false
	fmt.Println(name, age, isOk)

	// 批量声明变量
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)

	// 声明变量同时指定初始值
	var _name1 string = "shinrin"
	var _age1 int = 18
	fmt.Println(_name1, _age1)

	var _name2, _age2 = "sererin", 19
	fmt.Println(_name2, _age2)

	// 类型推导：编译器根据初始值推导类型
	var _name3 = "OHIUA"
	var _age3 = 20
	fmt.Println(_name3, _age3)

	// 短变量声明：只能在函数内使用
	m := 10
	fmt.Println(m)

	// 匿名变量：多值赋值时，使用匿名变量以忽略某个值。
	x, _ := fine()
	fmt.Println(x)

	// 注：函数外的每个语句必须以关键字开始。
}
