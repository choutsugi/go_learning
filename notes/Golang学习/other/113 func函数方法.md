## 普通函数定义

```go
func main() {
	r1, r2 := myFunc(0, "data2")
	fmt.Println(r1, r2)
}

func myFunc(data1 int, data2 string)(ret1 int, ret2 string) {
	ret1 = data1
	ret2 = data2
	return
}
```

## 函数类型

**函数类型：**将函数作为一种类型来定义变量（函数指针喵）。

```go
func myFunc(a, b int) (sum int) {
	sum = a + b
	return
}

// 定义函数类型（函数指针）
type FuncType func(a, b int) int

func main() {
	var result FuncType
	result = myFunc
	s := result(3, 5)
	fmt.Println(s)
}
```

## 匿名函数定义

```go
func main() {
	b := func(str string){
		fmt.Println(str)
	}

	b("这是一个匿名函数。")	// 调用匿名函数
}

```

以上等价于：

```go
func main() {
	func(str string){
		fmt.Println(str)
	}("这是一个匿名函数。")	// 定以后即调用
}
```

> 匿名函数的实质：在当前函数外定义一个函数，赋给函数内的一个变量（直接调用时省略该变量）。

## 不定参函数

```go
func main() {
	arr := []string{"8", "6", "6", "5", "1", "0"}
	zoe(1092, arr...)		// 将arr以不定项实参传入。
}

//不定参的实质：切片
func zoe(data1 int, data2 ...string){
	moe(data1, data2[1:]...)	// 传递不定参
}

func moe(data1 int, data2 ...string){
	fmt.Println(data1,data2)
	fmt.Printf("type of data2: %T\n", data2)
	fmt.Printf("data2: ")
	for _, v := range data2 {
		fmt.Printf("%s ", v)
	}
	fmt.Println("")
}
```

## 自执行函数

> 用于协程、定时任务。

```go
func main() {
	func(){
		fmt.Println("一个自执行函数。")
	}()	// 话说这玩意儿跟lambda好像。
}
```

## 闭包函数

即一个函数返回一个函数。

```go
func main() {
	str := mo()(7)
	fmt.Println(str)
}

func mo() func(num int) string {
	return func(num int) string {
		fmt.Println("这是一个闭包函数。")
		return "闭包函数返回值" + strconv.Itoa(num)
	}
}
```

## 延迟调用函数

***defer语句以FIFO的顺序在函数return之后执行。***

**应用场景：**文件、socket、数据库等的关闭。

```go
a, b := 10, 20
defer func(a, b int) {
    fmt.Println("unnamed func, a = ", a)
    fmt.Println("unnamed func, b = ", b)  
} (a, b)

a = 100
b = 200
fmt.Println("main func, a = ", a)
fmt.Println("main func, b = ", b)  
```

> defer虽然在return之后执行，但是其传参已在调用时完成。

## go函数

涉及异步、channel，后续学习。

