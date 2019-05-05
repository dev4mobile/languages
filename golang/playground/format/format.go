package format

import "fmt"

func Printf() {
	var protocol interface{}
	fmt.Printf("1. %T\n", protocol)
	fmt.Printf("2. %T\n", `hello`)
	fmt.Printf("3. %v\n", `hello`)
	fmt.Printf("4. %+v\n", `hello`)
	fmt.Printf("5. %#v\n", `hello`)
	fmt.Printf("6. %#v\n", "hello")
}
