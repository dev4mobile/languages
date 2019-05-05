package main

import (
	"fmt"
	"sync"
)

type ThreadSafeSet struct {
	sync.RWMutex
	s []int
}

func (set *ThreadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()
		fmt.Println("Rlock")
		for elem := range set.s {
			ch <- elem
		}
		close(ch)
		set.RUnlock()
		fmt.Println("Runlock")
	}()
	fmt.Println("Iter func return")
	return ch
}

func main() {
	set := ThreadSafeSet{}
	set.s = make([]int, 10)
	ch := set.Iter()
	closed := false
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println("read:", v)
			} else {
				closed = true
			}
		}
		if closed {
			fmt.Println("closed")
			break
		}
	}
	fmt.Println("done")
}
