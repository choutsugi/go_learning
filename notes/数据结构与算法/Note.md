## 算法

### 常用实例

#### 计算阶乘数列

```go
package main

import "fmt"

func Factorial(c chan int, n int) {
	ret := 1
	for i := 1; i <= n; i++ {
		ret = ret * i
		c <- ret
	}
	close(c)
}

func main() {
	c := make(chan int, 5)
	go Factorial(c, 5)
	for i := range c {
		fmt.Println(i)
	}
}
```

#### 计算斐波那契数列

```go
package main

import "fmt"

func Fibonacci(c chan int, n int) {

	a, b := 0, 1

	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
	}

	close(c)
}

func main() {
	c := make(chan int, 10)
	go Fibonacci(c, 5)
	for i := range c {
		fmt.Println(i)
	}
}
```

#### 二分查找

```go
package main

import "fmt"

func BinarySearch(array []int, target int, l, r int) int {

	ltemp := l
	rtemp := r

	for {
		if ltemp > rtemp {
			return -1
		}

		mid := (ltemp + rtemp) / 2
		midValue := array[mid]

		if target == midValue {
			return mid
		} else if target < midValue {
			rtemp = mid - 1
		} else {
			ltemp = mid + 1
		}
	}
}

func main() {
	array := []int{1, 5, 9, 15, 81, 123, 189, 333}
	target := 500
	i := BinarySearch(array, target, 0, len(array)-1)
	fmt.Println(i)

	target = 123
	i = BinarySearch(array, target, 0, len(array)-1)
	fmt.Println(i)

}
```

### 尾递归

递归函数调用自身后传回结果，不对其再加运算，提高了效率。

#### 尾递归计算阶乘

```go
package main

import (
	"fmt"
)

func FactorialTail(c chan int, n int, ret int) {
	if n == 1 {
		c <- ret
	}
	FactorialTail(c, n-1, ret*n)
}

func main() {
	c := make(chan int, 1)
	go FactorialTail(c, 5, 1)
	result, ok := <-c
	if ok {
		fmt.Println(result)	// 120
	}
}
```

#### 尾递归计算斐波那契数

```go
package main

import (
	"fmt"
)

func FibonacciTail(c chan int, n, a, b int) {
	if n == 1 {
		c <- a
	}
	FibonacciTail(c, n-1, b, a+b)
}

func main() {
	c := make(chan int, 1)
	go FibonacciTail(c, 5, 0, 1)
	result, ok := <-c
	if ok {
		fmt.Println(result)	// 3
	}
}
```

### 递归

#### 递归二分查找

```go
package main

import "fmt"

func BinarySearch(array []int, target int, l, r int) int {
	if l > r {
		return -1
	}

	mid := (l + r) / 2
	if target == array[mid] {
		return mid
	} else if target < array[mid] {
		return BinarySearch(array, target, l, mid-1)
	} else {
		return BinarySearch(array, target, mid+1, r)
	}
}

func main() {
	array := []int{1, 5, 9, 15, 81, 123, 189, 333}
	target := 500
	i := BinarySearch(array, target, 0, len(array)-1)
	fmt.Println(i)

	target = 123
	i = BinarySearch(array, target, 0, len(array)-1)
	fmt.Println(i)

}
```

## 数据结构

### 链表

#### 循环链表



