package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

// NewInt 自定义类型：通过type关键字
type NewInt int

// SayHello 为NewInt添加一个SayHello的方法
func (m NewInt) SayHello() {
	fmt.Println("Hello, 我是一个NewInt。")
}

// MyInt 类型别名：只在代码中存在，编译完成时不存在MyInt类型。
type MyInt = int

// 结构体
type person struct {
	name, city string
	age        int8
}

// newPerson 定义构造函数
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

// Dream person的方法
func (p person) Dream() {
	fmt.Printf("%s: so much money!\n", p.name)
}

// 指针类型的接受者
func (p *person) setAgeByRef(newAge int8) {
	p.age = newAge
}

// 值类型的接受者
func (p person) setAgeByValue(newAge int8) {
	p.age = newAge
}

//Person 结构体Person类型：匿名字段
type Person struct {
	string
	int
}

type test struct {
	a int8
	b int8
	c int8
	d int8
}

// Address 地址结构体
type Address struct {
	Province string
	City     string
}

// User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address // 嵌套结构体
}

// Animal 基类
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s：噔噔噔！\n", a.name)
}

// Dog 子类继承Animal基类
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) appel() {
	fmt.Printf("%s：汪汪汪~\n", d.name)
}

// Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

// Class 班级
type Class struct {
	Title    string
	Students []*Student
}

// Planet 行星
type Planet struct {
	ID   int    `json:"id"` // 指定json序列化时的key，否则默认使用字段名
	name string // 私有成员不能被json包访问
}

func main() {
	// 自定义类型与类型别名
	var a NewInt
	var b MyInt
	fmt.Printf("type of a: %T\n", a) // type of: main.NewInt
	fmt.Printf("type of b: %T\n", b) // type of b: int

	// 结构体实例化
	var p1 person
	p1.name = "HAKUNO"
	p1.city = "Xi'an"
	p1.age = 24

	fmt.Printf("p1 = %v\n", p1)  // p1 = {HAKUNO Xi'an 24}
	fmt.Printf("p1 = %#v\n", p1) // p1 = main.person{name:"HAKUNO", city:"Xi'an", age:24}

	// 匿名结构体：临时数据结构等场景使用。
	var hero struct {
		Name string
		Age  int
	}
	hero.Name = "Teemo"
	hero.Age = 7
	fmt.Printf("%#v\n", hero) // struct { Name string; Age int }{Name:"Teemo", Age:7}

	// 指针类型结构体
	var p2 = new(person)
	p2.name = "YaSuo"
	p2.city = "艾欧尼亚"
	p2.age = 35
	fmt.Printf("%T\n", p2)       // *main.person
	fmt.Printf("p2 = %#v\n", p2) // p2 = &main.person{name:"YaSuo", city:"艾欧尼亚", age:35}

	// 取结构体地址实例化
	p3 := &person{}
	fmt.Printf("%T\n", p3)     // *main.person
	fmt.Printf("p3=%#v\n", p3) // p3=&main.person{name:"", city:"", age:0}
	p3.name = "ZOE"
	p3.age = 19
	p3.city = "NULL"
	fmt.Printf("p3=%#v\n", p3) // p3=&main.person{name:"ZOE", city:"NULL", age:19}

	// 未初始化的结构体
	var p4 person
	fmt.Printf("p4=%#v\n", p4) // p4=main.person{name:"", city:"", age:0}

	// 使用键值初始化：不填写的字段默认为其零值。
	p5 := person{
		name: "pluie",
		city: "Xi'an",
		age:  16,
	}
	fmt.Printf("p5=%#v\n", p5) // p5=main.person{name:"pluie", city:"xi'an", age:16}

	// 使用值的列表初始化
	p6 := &person{
		"lettredamour",
		"Nanjing",
		17,
	}
	fmt.Printf("p6=%#v\n", p6) // p6=&main.person{name:"lettredamour", city:"Nanjing", age:17}

	// 结构体内存布局
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)

	// 空结构体不占用内存
	var v struct{}
	fmt.Println(unsafe.Sizeof(v)) // 0

	// 调用构造函数
	p7 := newPerson("damour", "Changsha", 24)
	fmt.Printf("%#v\n", p7) // &main.person{name:"damour", city:"Changsha", age:24}

	// 调用结构体的方法
	p7.Dream()

	// 指针类型的接受者
	p7.setAgeByRef(20)
	fmt.Printf("%#v\n", p7) // &main.person{name:"damour", city:"Changsha", age:20}

	// 值类型的接受者
	p7.setAgeByValue(21)
	fmt.Printf("%#v\n", p7) // &main.person{name:"damour", city:"Changsha", age:20}

	// 匿名字段
	p8 := Person{
		"rain",
		19,
	}
	fmt.Printf("%#v\n", p8) // main.Person{string:"rain", int:19}

	// 嵌套结构体
	user := User{
		Name:   "John",
		Gender: "Man",
		Address: Address{
			Province: "Hunan",
			City:     "Changsha",
		},
	}
	fmt.Printf("%#v\n", user) // main.User{Name:"John", Gender:"Man", Address:main.Address{Province:"Hunan", City:"Changsha"}}

	// 继承
	dog := &Dog{
		Feet: 4,
		Animal: &Animal{ // 嵌套的是结构体指针
			name: "美娜",
		},
	}
	dog.appel() //美娜：汪汪汪~
	dog.move()  //美娜：噔噔噔！

	// JSON序列化
	class := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}

	for i := 0; i < 10; i++ {
		stu := &Student{
			ID:     i,
			Gender: "Man",
			Name:   fmt.Sprintf("stu%02d", i),
		}
		class.Students = append(class.Students, stu)
	}

	// 序列化
	data, err := json.Marshal(class)
	if err != nil {
		fmt.Println("json marshal failed.")
		return
	}
	fmt.Printf("json:%s\n", data)
	// 反序列化
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	class = &Class{}
	err = json.Unmarshal([]byte(str), class)
	if err != nil {
		fmt.Println("json unmarshal failed.")
		return
	}
	fmt.Printf("%#v\n", class)

	// 结构体标签Tag
	planet := Planet{
		ID:   0,
		name: "Sun",
	}
	data, err = json.Marshal(planet)
	if err != nil {
		fmt.Println("json marshal failed.")
		return
	}
	fmt.Printf("json str: %s\n", data) // json str: {"id":0}

}
