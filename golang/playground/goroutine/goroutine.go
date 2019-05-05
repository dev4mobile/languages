package main

import (
	"time"
)

func main() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				arr[i] = i;
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
