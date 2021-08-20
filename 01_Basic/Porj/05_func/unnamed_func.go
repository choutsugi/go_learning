package main

import "fmt"

func main() {
	num := 9
	// 定义函数类型f绑定匿名函数
	f := func() {
		num++
		fmt.Println("this is a unnamed func, the global var num is ", num)
	}
	// 调用匿名函数执行
	f()
	fmt.Println("this is main func, the global var num is ", num)
}
