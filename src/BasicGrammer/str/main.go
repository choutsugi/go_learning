package main

import (
	"fmt"
	"strings"
)

func main() {
	// 求字符串长度
	str1 := "hello"
	fmt.Println(len(str1)) // 5
	str2 := "shinrin"
	fmt.Println(len(str2)) // 7

	// 字符串拼接
	fmt.Println(str1 + str2) // helloshinrin
	str3 := fmt.Sprintf("%s - %s", str1, str2)
	fmt.Println(str3) // hello - shinrin

	// 字符串分割
	str4 := "have a nice day!"
	fmt.Println(strings.Split(str4, " "))        // [have a nice day!]
	fmt.Printf("%T\n", strings.Split(str4, " ")) // []string

	// 判断是否包含字串
	fmt.Println(strings.Contains(str4, "nice")) // true

	// 判断前缀
	fmt.Println(strings.HasPrefix(str4, "have")) // true

	// 判断后缀
	fmt.Println(strings.HasSuffix(str4, "day")) // false

	// 判断子串的位置
	fmt.Println(strings.Index(str4, "a")) // 1

	// 最后子串出现的位置
	fmt.Println(strings.LastIndex(str4, "a")) // 13

	// join
	str5 := []string{"have", "a", "nice", "day", "!"}
	fmt.Println(str5)                    // [have a nice day !]
	fmt.Println(strings.Join(str5, "-")) // have-a-nice-day-!
}
