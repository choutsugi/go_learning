package Common

func Fibonacci(c chan int, n int) {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		if i == n-1 {
			c <- a
		}
		a, b = b, a+b
	}
}
