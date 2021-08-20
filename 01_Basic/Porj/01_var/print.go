package main

import "fmt"

func main() {
	a, b, c := 1, 2, 3

	// fmt.Println
	fmt.Println("a = ", a, " b = ", b, " c = ", c)
	fmt.Printf("\n")

	// fmt.Print
	fmt.Print("a = ", a, " b = ", b, " c = ", c)
	fmt.Printf("\n")

	// fmt.Printf
	fmt.Printf("a = %d, b = %d, c = %d", a, b, c)
	fmt.Printf("\n")
}
