/*
	指针：同C++。
*/

package main

import "fmt"

func changeValue(val *int) {
	*val = 100
}

// 交换数值。
func swap(pa, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	var val int = 10
	changeValue(&val)
	fmt.Println("a = ", val)

	a, b := 20, 30
	fmt.Println("before swap: a = ", a, ", b = ", b)
	swap(&a, &b)
	fmt.Println("after swap: a = ", a, ", b = ", b)
}
