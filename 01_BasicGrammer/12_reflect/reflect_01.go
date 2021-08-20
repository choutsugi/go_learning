package main

import (
	"fmt"
	"reflect"
)

// // 利用反射获取变量的类型和值
// func reflectNum(arg interface{}) {
// 	fmt.Println("type: ", reflect.TypeOf(arg))
// 	fmt.Println("value: ", reflect.ValueOf(arg))

// }

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", this)
}

func DoFiledAndMethod(input interface{}) {
	// 获取type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is: ", inputType.Name())

	// 获取value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is: ", inputValue)

	// // 通过type获取字段
	// for i := 0; i < inputType.NumField(); i++ {
	// 	field := inputType.Field(i)
	// 	value := inputValue.Field(i).Interface()

	// 	fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	// }

	// // 通过type获取方法
	// for i := 0; i < inputType.NumMethod(); i++ {
	// 	m := inputType.Method(i)
	// 	fmt.Printf("%s: %v\n", m.Name, m.Type)
	// }

	for i := 0; i < inputValue.Elem().NumField(); i++ {
		eleValue := inputValue.Elem().Field(i)
		fmt.Println("element ", i, " its type is ", eleValue.Type())
		fmt.Println("element ", i, " its values is ", eleValue)
	}
}

func main() {
	// var num float64 = 1.2345
	// reflectNum(num)

	user := User{1, "Amo", 17}
	// user.Call()
	DoFiledAndMethod(user)

}
