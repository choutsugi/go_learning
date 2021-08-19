// 每个go文件必须归属于一个package，包含main函数的文件属于main包。
package main

// 通过import引入一个或多个package。
import (
	"fmt"
	"time"
)

func goFunc(i int) {
	// 调用fmt包中的Println函数。
	fmt.Println("goroutine ", i)
}

func main() {
	for i := 0; i < 1000; i++ {
		go goFunc(i) //开启协程，原生支持并发
	}
	// 调用time包中的Sleep函数。
	time.Sleep(time.Second)
}
