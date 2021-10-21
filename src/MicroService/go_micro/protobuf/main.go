package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	. "go_micro/protobuf/test"
)

func main() {
	// 学生组
	group := &Group{}

	// 学生1
	stu1 := &Student{}
	stu1.Name = "OHIUA"
	stu1.Age = 23
	stu1.Address = "JingChuan"
	stu1.Cn = ClassName_class1

	// 学生2
	stu2 := &Student{}
	stu2.Name = "HAKUNO"
	stu2.Age = 22
	stu2.Address = "DuanZhou"
	stu2.Cn = ClassName_class2

	group.Person = append(group.Person, stu1)
	group.Person = append(group.Person, stu2)
	group.School = "LUT"

	fmt.Printf("原始数据：%v\n", group)

	// 编码（序列化）
	buffer, _ := proto.Marshal(group)
	fmt.Printf("序列化：\t%v\n", buffer)

	// 解码（反序列化）
	data := &Group{}
	err := proto.Unmarshal(buffer, data)
	if err != nil {
		fmt.Printf("err:%v\n", err.Error())
		return
	}
	fmt.Printf("反序列化：%v\n", data)
}
