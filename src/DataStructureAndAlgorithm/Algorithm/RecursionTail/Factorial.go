package RecursionTail

func FactrialTail(c chan int64, n int, ret int64) {
	if n == 1 {
		c <- ret
	}

	FactrialTail(c, n-1, ret*int64(n))
}
