package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", v.Name(), v.Kind())
}

// 通过反射获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 通过反射设置值
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// 使用Elem，否则panic
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

// 结构体反射示例
type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (s student) Study() string {
	msg := "Study"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "Sleep"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func main() {
	var a float32 = 3.14
	reflectType(a) // type:float32 kind:float32
	var b myInt = 1997
	reflectType(b) // type:myInt kind:int64

	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b) // type is int64, value is 1997

	// 将原始类型int转换为reflect.Value类型。
	c := reflect.ValueOf(10)
	fmt.Printf("type c: %T\n", c) // type c: reflect.Value

	// 通过反射设置值
	reflectSetValue(&a)
	fmt.Println(a)

	// isNil 和 isValid
	// *int类型空指针
	var d *int
	fmt.Println("var d *int IsNil:", reflect.ValueOf(d).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化匿名结构体
	e := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(e).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(e).MethodByName("abc").IsValid())
	// map
	f := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(f).MapIndex(reflect.ValueOf("stone")).IsValid())

	// 结构体反射
	stu := student{
		Name:  "stone",
		Score: 95,
	}
	t := reflect.TypeOf(stu)
	fmt.Println(t.Name(), t.Kind())
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}

	// 通过反射调用方法
	printMethod(stu)
}
