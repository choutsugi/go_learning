package main

import "fmt"

// 常量声明
const pi1 = 3.1415
const e1 = 2.7

// 常量批量声明
const (
	pi2 = 3.1415
	e2  = 2.7
)

// 常量批量声明默认值
const (
	n1 = 10 // 10
	n2      // 10
	n3      // 10
)

// iota：常量计数器
const (
	m1 = iota // 0
	m2        // 1
	_         // 2，使用_跳过
	m4 = 100  // 100，插入100代替3
	m5        // 100
	m6 = iota // 5
)

// iota：遇const重置为0
const m7 = iota // 0

// iota：定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) // 1<<10
	MB = 1 << (10 * iota) // 1<<20
	GB = 1 << (10 * iota) // 1<<30
	TB = 1 << (10 * iota) // 1<<30
	PB = 1 << (10 * iota) // 1<<40
)

// iota：多个iota定义在同一行
const (
	a, b = iota + 1, iota + 2 // iota = 0, 1,2
	c, d                      // iota = 1, 2,3
	e, f                      // iota = 2, 3,4
)

func main() {
	fmt.Println(pi1, e1)
	fmt.Println(pi2, e2)
	fmt.Println(n1, n2, n3)
	fmt.Println(m1, m2, m4, m5, m6, m7)
	fmt.Println(KB, MB, GB, TB, PB)
	fmt.Println(a, b, c, d, e, f)
}
