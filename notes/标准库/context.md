# context

## 使用场景

### 退出协程

方式一：通过全局变量。

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
	exit bool
)

func worker() {
	defer wg.Done()
	for{
		fmt.Println("worker...")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go worker()
	time.Sleep(time.Second * 5)
	exit = true
	wg.Wait()
	fmt.Println("over...")
}
```

方式二：通过channel

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func worker(ch <- chan struct{}) {
	defer wg.Done()
	LABEL:
	for{
		select {
		case <-ch:
			break LABEL
		default:
			fmt.Println("worker...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var exitChan = make(chan struct{}, 1)
	wg.Add(1)
	go worker(exitChan)
	time.Sleep(time.Second * 5)
	exitChan <- struct{}{}
	wg.Wait()
	fmt.Println("over...")
}
```

方式三：通过context

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

// context内部封装了channel
func worker(ctx context.Context) {
	defer wg.Done()
	LABEL:
	for{
		fmt.Println("worker...")
		time.Sleep(time.Second)
		select {
		case <- ctx.Done():
			break LABEL
		default:
			//...
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over...")
}
```

## 简介

在Go http的Server中，请求处理函数启动额外`goroutine`访问后端服务，当一个请求被取消或超时时，所有用于处理该请求的`goroutine`都应退出；`context`中的`Context`类型用于简化请求域数据、取消信号、截止时间等相关操作。

### Context接口

#### Background()和TODO()

`Background()`：用于`main`函数的根`context`。

`TODO()`：未知场景使用。

#### WithCancel()

返回带有`Done`通道的父节点副本和`CancelFunc`，调用`CancelFunc()`关闭`Done通道`，取消此上下文将释放与其关联的资源。

```go
package main

import (
	"context"
	"fmt"
)

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
```

#### WithDeadline()

示例：

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 过期时间：50ms
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		// 不会被执行
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
```

#### WithTimeout()

用于数据库或网络连接的超时控制。

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("db connecting...")
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:
			// nothing
		}
	}
	fmt.Println("worker done")
}

func main() {
	// 超时时间：50ms
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)

	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
```

#### WithValue()

使用`context`传递值，为避免冲突，应使用自定义类型。

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type UserID string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	defer wg.Done()

	key := UserID("user_id")
	traceCode, ok := ctx.Value(key).(int64)
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code: %d\n", traceCode)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, UserID("user_id"), int64(123456))
	wg.Add(1)
	go worker(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
```



## 知识点

***make和new的区别***

相同点：都用来初始化内存。

区别：

- new多用于基本数据类型（bool、string、int）的初始化内存，返回指针。

- make用于初始化`slice`、`map`、`channel`返回对应的类型。

***channel做标识使用空结构体***

特点：

- 空结构体，不占用内存（宽度为0），适用于占位符场景（无需发送数据，只作为通知使用）。
- 自身即零值。