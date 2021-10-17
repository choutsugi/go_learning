package main

import "fmt"

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

// 计数
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

// 平方
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

// 打印
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	// 无缓冲channel：对端接收时才能写入。
	ch1 := make(chan int)
	go recv(ch1) // 启用goroutine从通道接收值
	ch1 <- 10
	fmt.Println("[无缓冲channel]发送成功")

	// 有缓冲channel
	ch2 := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch2 <- 10
	fmt.Println("[有缓存channel]发送成功")

	// for range循环取值
	ch3 := make(chan int)
	ch4 := make(chan int)
	// 开启goroutine将0~100的数发送到ch3中
	go func() {
		for i := 0; i < 100; i++ {
			ch3 <- i
		}
		close(ch3)
	}()
	// 开启goroutine从ch3中接收值，并将该值的平方发送到ch4中
	go func() {
		for {
			i, ok := <-ch3 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch4 <- i * i
		}
		close(ch4)
	}()
	// 在主goroutine中从ch4中接收值打印
	for i := range ch4 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}

	// 单向通道
	ch5 := make(chan int)
	ch6 := make(chan int)
	go counter(ch5)
	go squarer(ch6, ch5)
	printer(ch6)
}
