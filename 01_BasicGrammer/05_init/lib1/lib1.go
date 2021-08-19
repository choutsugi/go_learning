package lib1

import "fmt"

// 函数名首字母大写：对外开放，包外可访问。
func Lib1Test() {
	fmt.Println("lib1.lib2Test()...")
}

// Lib1的初始化函数：对外隐藏，包外不可访问。
func init() {
	fmt.Println("lib1.init()...")
}
