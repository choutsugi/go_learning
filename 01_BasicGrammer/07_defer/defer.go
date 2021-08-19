package main

import "fmt"

func func1() {
	fmt.Println("defer func1 ..")
}

func func2() {
	fmt.Println("defer func2 ..")
}

func func3() {
	fmt.Println("defer func3 ..")
}

func deferFunc() int {
	fmt.Println("defer function called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return function called...")
	return 0
}

// 判断defer与return的执行顺序
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	// defer：程序结束时执行，多个defer以栈的方式存入（先定义的后执行）
	// 以下执行顺序：func3、func2、func1
	defer func1()
	defer func2()
	defer func3()

	// defer 与 return：return先执行、defer后执行。
	returnAndDefer()

}
