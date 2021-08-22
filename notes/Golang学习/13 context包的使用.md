## Context上下文

**原始方法：**通过管道数据控制协程。

```go
func main() {

	flag := make(chan bool)
	message := make(chan int)

	go son(flag, message)
	for i := 0; i < 10; i++ {
		message <- i
	}
	flag <- true
	time.Sleep(time.Second)
	fmt.Println("main() end.")
}

func son(flag chan bool, message chan int){
	t := time.Tick(time.Second)
	for _ = range t {	// 每秒钟执行一次
		select {
		case m := <- message:
			fmt.Printf("接受数据：%d\n", m)
		case <- flag:
			fmt.Println("son() end.")
			return
		}
	}
}
```

***WithCancel()***

```go
func main() {
	// 携带参数
	ctx := context.WithValue(context.Background(), "name", "ohiua")
	ctx, clear := context.WithCancel(ctx)

	message := make(chan int)

	go son(ctx, message)
	for i := 0; i < 10; i++ {
		message <- i
	}
	clear()
	time.Sleep(2 * time.Second)
	fmt.Println("main() end.")
}

func son(ctx context.Context, message chan int){
	t := time.Tick(time.Second)
	for _ = range t {	// 每秒钟执行一次
		select {
		case m := <- message:
			fmt.Printf("接受数据：%d\n", m)
		case <- ctx.Done():
			fmt.Println("son() end.", ctx.Value("name"))
			return
		}
	}
}
```

***WithDeadline()***

```go
func main() {

	ctx, clear := context.WithDeadline(context.Background(), time.Now().Add(time.Second * 5))
	message := make(chan int)

	go son(ctx, message)
	for i := 0; i < 10; i++ {
		message <- i
	}
	defer clear()
	time.Sleep(2 * time.Second)
	fmt.Println("main() end.")
}

func son(ctx context.Context, message chan int){
	t := time.Tick(time.Second)
	for _ = range t {	// 每秒钟执行一次
		select {
		case m := <- message:
			fmt.Printf("接受数据：%d\n", m)
		case <- ctx.Done():
			fmt.Println("son() end.")
			return
		}
	}
}
```

***WithTimeout()***

```go
func main() {

	ctx, clear := context.WithTimeout(context.Background(), time.Second * 4)
	message := make(chan int)

	go son(ctx, message)
	for i := 0; i < 10; i++ {
		message <- i
	}
	defer clear()
	time.Sleep(2 * time.Second)
	fmt.Println("main() end.")
}

func son(ctx context.Context, message chan int){
	t := time.Tick(time.Second)
	for _ = range t {	// 每秒钟执行一次
		select {
		case m := <- message:
			fmt.Printf("接受数据：%d\n", m)
		case <- ctx.Done():
			fmt.Println("son() end.")
			return
		}
	}
}
```

**传递context**

```go
func main() {

	ctx, clear := context.WithTimeout(context.Background(), time.Second * 4)
	message := make(chan int)

	go son(ctx, message)
	for i := 0; i < 10; i++ {
		message <- i
	}
	defer clear()
	time.Sleep(2 * time.Second)
	fmt.Println("main() end.")
}

func son(ctx context.Context, message chan int){
	t := time.Tick(time.Second)
	ctxT, clean := context.WithCancel(ctx)
	go son(ctxT, message)
	defer clean()
	for _ = range t {	// 每秒钟执行一次
		select {
		case m := <- message:
			fmt.Printf("接受数据：%d\n", m)
		case <- ctx.Done():
			fmt.Println("son() end.")
			return
		}
	}
}
```



