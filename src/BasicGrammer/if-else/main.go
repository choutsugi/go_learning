package main

import "fmt"

// if判断
func main() {
	// if判断基本写法
	var scoreA = 65
	if scoreA >= 90 {
		fmt.Println("A")
	} else if scoreA > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// if判断的特殊写法
	if scoreB := 65; scoreB >= 90 {
		fmt.Println("A")
	} else if scoreB > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
