// Package pkg 定义包
package pkg

// 导入包
import (
	"fmt"
)

// 包变量的可见性：通过标识符首字母大小写区分

// Mode 可在包外访问的变量
const Mode = 1

// person 仅限包内访问的结构体
type person struct {
	name string // 仅限包内访问的结构体成员
}

// Add 可在包外访问的方法
func Add(x, y int) int {
	return x + y
}

// age 仅限包内访问的方法
func age() {
	var Age = 18
	fmt.Println(Age)
}
