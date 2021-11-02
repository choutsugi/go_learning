package RecursionTail

func FibonacciTail(c chan int, n int, a, b int) {
	if n == 1 {
		c <- a
	}

	FibonacciTail(c, n-1, b, a+b)
}
