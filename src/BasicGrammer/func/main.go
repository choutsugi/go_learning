package main

import (
	"errors"
	"fmt"
	"strings"
)

// 无参无返回值函数
func sayHello() {
	fmt.Println("hello ohiua")
}

// 有参无返回值函数
func sayMoring(name string) {
	fmt.Printf("Good moring, Mr.%s\n", name)
}

// 有参带返回值函数
func addNum(a int, b int) (sum int) {
	sum = a + b
	return
}

// 变参函数：可变参数在参数列表的最后
func doSum(val int, args ...int) (name string, ret int) {
	fmt.Println(args)
	name = "let's"
	ret = val
	for _, arg := range args {
		ret = ret + arg
	}
	return
}

// 函数类型
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 函数作为参数
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("Undefined operator" + s)
		return nil, err
	}
}

// 闭包：闭包 = 函数+引用函数
// 闭包示例1：
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包示例2：
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	// 普通函数调用
	name := "HAKUNO"
	sayHello()
	sayMoring(name)
	sum := addNum(3, 2)
	fmt.Println("sum of 3 and 2 is", sum)

	// 变参函数调用
	who, ret := doSum(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(who, "doSum:", ret)

	// 函数类型变量
	var cal calculation                 // 声明一个calculation类型的变量cal
	cal = add                           // 把add赋值给cal
	fmt.Printf("type of cal:%T\n", cal) // type of cal:main.calculation
	fmt.Println(cal(1, 2))              // 3

	fun := sub                          // 将函数sub赋值给变量fun
	fmt.Printf("type of fun:%T\n", fun) // type of fun:func(int, int) int
	fmt.Println(fun(10, 20))            // -10

	// 函数作为参数
	retVal1 := calc(10, 20, add) // 传入add函数
	fmt.Println(retVal1)         // 30

	// 函数作为返回值
	retVal2, _ := do("+")        // 返回add函数
	fmt.Println(retVal2(30, 20)) // 50

	// 匿名函数：匿名函数可在函数内定义，需要保存到某个变量或者立即执行；多用于实现回调函数或闭包。
	mul := func(x, y int) {
		fmt.Println("mul: ", x*y)
	}
	mul(2, 10) // 调用匿名函数

	// 自执行函数：匿名函数定义完成后加()立即执行。
	func(x, y int) {
		fmt.Println(x / y)
	}(50, 10)

	// 闭包示例1：
	var fun1 = adder()
	fmt.Println(fun1(10)) // 10
	fmt.Println(fun1(20)) // 30
	fmt.Println(fun1(30)) // 60

	fun2 := adder()
	fmt.Println(fun2(40)) // 40
	fmt.Println(fun2(50)) // 90

	// 闭包示例2：
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt

	// defer
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}
