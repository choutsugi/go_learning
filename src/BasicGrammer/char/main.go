package main

import "fmt"

// 字符
func main() {
	// byte uint8别名 ASCII码
	// rune int32别名

	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)                 // 99 99
	fmt.Printf("c1:%T c2:%T\n", c1, c2) // c1:uint8 c2:int32

	str := "hello 白野桑"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i]) // 按字节输出：中文乱码
	}

	for _, v := range str {
		fmt.Printf("%c\n", v) // 按字符输出：中文正常
	}
}
