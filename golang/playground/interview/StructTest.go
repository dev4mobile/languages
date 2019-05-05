package main

import "fmt"

type Speaker interface {
	Speak(string) string
}

type Student struct {
}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "you are a good body"
	} else {
		talk = "hi"
	}
	return talk
}

func main() {
	var speaker Speaker = &Student{}
	fmt.Printf("%T\n", speaker)
	fmt.Printf("%V\n", speaker)
	fmt.Printf("%v\n", speaker)
}

