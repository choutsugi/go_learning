package main

import (
	"fmt"
)

type Hero struct {
	Name  string
	Hp    int
	level int // 小写：private属性
}

// 大写：public方法
func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Hp = ", this.Hp)
	fmt.Println("Level = ", this.level)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	hero := Hero{Name: "Teemo", Hp: 700, level: 1}
	hero.Show()
	hero.SetName("Yasuo")
	name := hero.GetName()
	fmt.Println(name)
}
