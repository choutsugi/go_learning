package main

import (
	"fmt"
)

// 接口：本质是指针
type AnimalIF interface {
	Sleep()
	GetColor() string // 颜色
	GetType() string  // 种类
}

// 接口实现类：Cat
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Printf("Cat is sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 接口实现类：Dog
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

func main() {
	/*
		var animal AnimalIF // 接口类型：父类指针
		animal = &Cat{"Blue"}
		animal.Sleep()
		animal = &Dog{"Red"}
		animal.Sleep()
	*/

	cat := Cat{"Blue"}
	dog := Dog{"Red"}

	showAnimal(&cat)
	showAnimal(&dog)
}
