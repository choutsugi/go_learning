package main

import "fmt"

func main() {

	var str string
	str = "abc"
	fmt.Println("str = ", str)
	fmt.Printf("str = %s", str)

	fmt.Println("the length of str is ", len(str))
	fmt.Printf("str[0] = %c, str[1] = %c", str[0], str[1])
}
