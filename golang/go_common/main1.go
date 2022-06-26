package main

import (
	"fmt"
)

func main() {
	m := [...]int{
		'^': 1,
		'a': 2,
		'b': 3,
		'z': 4,
		'Z': 20,
	}
	// m['a'] = 10
	// fmt.Println(m)
	fmt.Println(len(m))
	for i, v := range m {
		// fmt.Print(i, string(v))
		fmt.Print(i, string(v))
		fmt.Print(",")
	}
}
