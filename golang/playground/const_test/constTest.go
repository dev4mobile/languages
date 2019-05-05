package const_test

import "fmt"

func consts() {
	const fileName = "abc.txt"
	//const a = 1
	//const b = 1
	//var c = math.Sqrt(a*a + b*b)
	//fmt.Println(c)
	//const (
	//	a1 = 35
	//	b1 = 40
	//)
	const(
		a = iota
		b
		c = iota
		d
	)
	fmt.Errorf("%d",3)
}
