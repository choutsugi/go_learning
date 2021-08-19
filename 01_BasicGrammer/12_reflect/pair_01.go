package main

import (
	"fmt"
)

func main() {
	var s string
	// pair<statictype:string, value:"aceld">
	s = "acdld"
	// pair<type:string, value:"aceld">

	var allType interface{}

	allType = s

	str, _ := allType.(string)
	fmt.Println(str)

}
