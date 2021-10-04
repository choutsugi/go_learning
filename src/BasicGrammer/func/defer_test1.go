package main

import "fmt"

// defer中x++，但未返回
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// defer中x++，并返回
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

// 返回x到y
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
}
