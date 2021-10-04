## Mutex互斥锁

**互斥锁：**对代码块加锁，使得同一时刻只有一个协程可以访问。

```go
func SyncClass(){
	l := &sync.Mutex{}

	go lockFunc(l)
	go lockFunc(l)
	go lockFunc(l)
	go lockFunc(l)
	go lockFunc(l)

	time.Sleep(5 * time.Second)
}

func lockFunc(lock *sync.Mutex){
	lock.Lock()
    defer lock.Unlock()
	fmt.Println(" ohiua")
	time.Sleep(1 * time.Second)
}
```

## RWMutex读写互斥锁

**写锁：**阻塞所有读写操作。

**读锁：**阻塞所有写操作。

```go
func SyncClass(){
	l := &sync.RWMutex{}

	go lockFunc(l)
	go lockFunc(l)
	go readLockFunc(l)
	go readLockFunc(l)
	go readLockFunc(l)
	go readLockFunc(l)

	time.Sleep(5 * time.Second)
}

func lockFunc(lock *sync.RWMutex){
	lock.Lock()		// 写锁（互斥锁）：阻塞所有读写。
    defer lock.Unlock()
	fmt.Println(" ohiua RWLock")
	time.Sleep(1 * time.Second)
}

func readLockFunc(lock *sync.RWMutex){
	lock.RLock()	// 读锁：阻塞写操作，禁止写入。
    defer lock.RUnlock()
	fmt.Println(" ohiua RLock")
	time.Sleep(1 * time.Second)
}
```

## Once

仅执行一次。

```go
func SyncClass(){
	oneDo := &sync.Once{}
	for i := 0; i < 10; i++ {
		oneDo.Do(func(){
			fmt.Println(i)
		})
	}
}
```

## WaitGroup

等待组：等待协程结束后再结束主线程。

```go
func SyncClass(){
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(){
		time.Sleep(8 * time.Second)
		fmt.Println("goroutine end with 8 second go by.")
		wg.Done()
	}()

	go func(){
		time.Sleep(6 * time.Second)
		fmt.Println("goroutine end with 6 second go by.")
		wg.Done()
	}()

	wg.Wait()
}
```

## 并发字典

多协程可读写的map：

```go
func SyncClass(){
	m := &sync.Map{}

	go func(){
		for{
			m.Store(1, 1)
		}
	}()

	go func(){
		for{
			v, flag := m.Load(1)
			if flag {
				fmt.Println("读取数据：", v)
			}
		}
	}()

	time.Sleep(100)
}
```

其他操作：

```go
func SyncClass(){
	m := &sync.Map{}
	m.Store(1,1)
	m.Store(2,2)
	m.Store(3,3)
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key,value)
		time.Sleep(1 * time.Second)
		return true
	})
}
```

## Pool并发池

存放数据供所有协程使用？

```go
func SyncClass(){
	pool := &sync.Pool{}

	for i := 0; i < 6; i++ {
		pool.Put(i)
	}

	for i := 0; i < 8; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(pool.Get())
	}
}
```

## Cond条件变量

```go
func SyncClass(){
	cond := sync.NewCond(&sync.Mutex{})
	flag := false

	go func() {
		cond.L.Lock()
		fmt.Println("lock 1")
		for flag {
			cond.Wait()
		}
		fmt.Println("unlock 1")
		cond.L.Unlock()
	}()

	go func() {
		cond.L.Lock()
		fmt.Println("lock 2")
		for flag {
			cond.Wait()
		}
		fmt.Println("unlock 2")
		cond.L.Unlock()
	}()

	//time.Sleep(2 * time.Second)
	//cond.Broadcast()	// 唤醒全部等待中的协程
	//time.Sleep(1 * time.Second)

	time.Sleep(2 * time.Second)
	cond.Signal()		// 唤醒一个等待中的协程
	time.Sleep(1 * time.Second)
	cond.Signal()		// 再次唤醒一个等待中的协程
	time.Sleep(1 * time.Second)
}
```

