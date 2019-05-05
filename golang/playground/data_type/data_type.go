package data_type

import "fmt"

func DataType() {
	var a interface{}
	a = 5
	//a = "hello world"
	fmt.Printf("a's type: %T\n", a)

	v, ok := a.(string)
	fmt.Printf("value=%v, ok=%v\n", v, ok)
}
