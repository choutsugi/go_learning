package lib2

import "fmt"

// 函数名首字母大写：对外开放
func Lib2Test() {
	fmt.Println("lib2.lib2Test()...")
}

// Lib2的初始化函数
func init() {
	fmt.Println("lib2.init()...")
}
