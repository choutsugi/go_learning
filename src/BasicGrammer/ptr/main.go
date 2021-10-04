package main

import "fmt"

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

// 指针
func main() {
	// 指针地址和指针类型
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00000a088
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00000a088 type:*int
	fmt.Println(&b)                    // 0xc000006028

	// 指针取值（根据指针去内存取值）
	fmt.Printf("type of b:%T\n", b) // type of b:*int
	c := *b
	fmt.Printf("type of c:%T\n", c)  // type of c:int
	fmt.Printf("value of c:%v\n", c) // value of c:10

	// 指针传值
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100

	// new
	d := new(int)
	e := new(bool)
	fmt.Printf("type of d: %T\n", d) // type of d: *int
	fmt.Printf("type of e: %T\n", e) // type of e: *bool
	fmt.Println(*d)                  // 0
	fmt.Println(*e)                  // false

	// 指针变量初始化后才能赋值
	var f *int
	f = new(int)
	*f = 10012
	fmt.Println(*f) // 10012

	var g map[string]int
	g = make(map[string]int, 10)
	g["Hasaki"] = 100
	fmt.Println(g) // map[Hasaki:100]
}
