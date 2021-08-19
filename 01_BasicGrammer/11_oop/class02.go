package main

import (
	"fmt"
)

type Human struct {
	name string
	age  int
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human // 继承Human
	level int
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("age = ", this.age)
	fmt.Println("level = ", this.level)
}

func main() {
	h := Human{"Mengduo", 22}
	h.Eat()
	h.Walk()

	// s := SuperMan{Human{"Hasaki", 22}, 12}
	var s SuperMan
	s.name = "Hasaki"
	s.age = 22
	s.level = 12

	s.Eat()
	s.Walk()
	s.Fly()
	s.Print()
}
