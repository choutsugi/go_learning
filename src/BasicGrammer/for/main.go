package main

import "fmt"

func main() {
	// for循环基本写法
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 省略初始语句，保留分号
	var j = 0
	for ; j < 10; j++ {
		fmt.Println(j)
	}

	// 省略初始语句和结束语句
	var k = 10
	for k > 0 {
		fmt.Println(k)
		k--
	}

	// 死循环
	// for {
	// 	fmt.Println("int the loop")
	// }

	// break跳出for循环
	for l := 0; l < 5; l++ {
		if l == 3 {
			break
		}
		fmt.Println(l)
	}

	// continue跳过本次循环，继续下次循环
	for m := 0; m < 5; m++ {
		if m == 3 {
			break
		}
		fmt.Println(m)
	}

	// for range 用于遍历数组、切片、字符串、map和channel。
	slice := []string{"have", "a", "nice", "day", "!"}
	for _, v := range slice {
		fmt.Println(v)
	}
}
