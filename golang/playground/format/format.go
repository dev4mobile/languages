package format

import "fmt"

func Printf() {
	//var protocol interface{}
	//fmt.Printf("1. %T\n", protocol)
	//fmt.Printf("2. %T\n", `hello`)
	//fmt.Printf("3. %v\n", `hello`)
	//fmt.Printf("4. %+v\n", `hello`)
	//fmt.Printf("5. %#v\n", `hello`)
	//fmt.Printf("6. %#v\n", "hello")
	var i int
	var p *int
	var ptr **int
	p = &i
	ptr = &p
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", ptr)
}

type foo struct{}

func (f *foo) hello() {
	fmt.Println("f hello")
}

type bar struct{}

func (b *bar) hello() {
	fmt.Println("b hello")
	(*foo)(b).hello()
	a := "helloworld"
	c := []rune(a)
	fmt.Printf("%T", c)
}
