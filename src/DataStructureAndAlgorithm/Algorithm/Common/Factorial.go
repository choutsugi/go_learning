package Common

func Factrial(c chan int64, n int) {
	var ret int64 = 1
	for i := 1; i <= n; i++ {
		ret = ret * int64(i)
		if i == n {
			c <- ret
		}
	}
}
