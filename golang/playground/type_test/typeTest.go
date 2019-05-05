package type_test

import "fmt"

func Type() {
	a := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Printf("I'm bool.")
		case int:
			fmt.Printf("I'm an int.")
		default:
			fmt.Printf("hello")
		}
	}

	a(3)
}
