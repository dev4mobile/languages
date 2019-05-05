package ops_channel

import "fmt"

var ch chan int = make(chan int)

func SamplChannel() {
	fmt.Printf("go start")
	<-ch
}
