package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu Student) Speak(think string) (talk string) {
	if think == "LOL" {
		talk = "hhh.."
	} else {
		talk = "Hasaki"
	}
	return
}

func main() {
	var peo People = Student{}
	think := "lol"
	fmt.Println(peo.Speak(think))
}
