package main

import "fmt"

func main() {

	const (
		a = iota // 0
		b        // 1
		c        // 2
	)
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	const (
		i          = iota
		j1, j2, j3 = iota, iota, iota
		k          = iota
	)
	fmt.Printf("i = %d, j1 = %d, j2 = %d, j3 = %d, k = %d\n", i, j1, j2, j3, k)

	const (
		m, n = iota + 1, iota + 2 // iota = 0, m = 1, n = 2
		o, p                      // iota = 1, o = 2, p = 3
		q, r                      // iota = 2, q = 3, r = 4
		s, t = iota * 2, iota * 3 // iota = 3, s = 6, t = 9
		u, v                      // iota = 4, u = 8, v = 12
	)
}
