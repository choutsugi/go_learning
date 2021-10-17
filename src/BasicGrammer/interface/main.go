package main

import "fmt"

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

// 结构体
type dog struct{}

type cat struct{}

// 实现接口
func (d dog) say() {
	fmt.Println("嗷呜！")
}

func (c cat) say() {
	fmt.Println("喵呜~")
}

// 值接收者实现接口
func (d dog) move() {
	fmt.Println("狗刨！")
}

// 指针接收者实现接口
func (c *cat) move() {
	fmt.Println("猫抓！")
}

// 空接口应用：作为函数参数。
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	// 接口类型变量
	var s Sayer

	a := dog{}
	b := cat{}

	s = a
	s.say()

	s = b
	s.say()

	// 值接收者实现接口
	var m Mover
	var mq = dog{}
	m = mq // 可接受dog类型
	var qm = &dog{}
	m = qm // 可接受*dog类型
	m.move()

	// 指针接收者实现接口
	//var mn = cat{}
	//m = mn // 不可接收cat类型
	var nm = &cat{}
	m = nm
	m.move()

	// 接口嵌套
	var anl animal
	anl = dog{}
	anl.say()
	anl.move()

	// 空接口
	var itf interface{}
	val := "hello stone"
	itf = val
	fmt.Printf("type:%T value:%v\n", itf, itf)
	i := 100
	itf = i
	fmt.Printf("type:%T value:%v\n", itf, itf)
	bool := true
	itf = bool
	fmt.Printf("type:%T value:%v\n", itf, itf)

	// 空接口应用：作为map值。
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "stone"
	studentInfo["age"] = 18
	fmt.Println(studentInfo)

	// 类型断言
	v, ok := itf.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
