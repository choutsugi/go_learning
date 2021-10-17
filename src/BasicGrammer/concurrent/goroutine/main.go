package main

import (
	"fmt"
	"sync"
	"time"
)

func hello() {
	fmt.Println("Hello Goroutine!")
}

var wg sync.WaitGroup

func helloMul(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}

func main() {
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second) // 等待hello执行完成

	for i := 0; i < 10; i++ {
		wg.Add(1) // 每启动一个goroutine登记一次
		go helloMul(i)
	}

	wg.Wait()
}
