package read_mutex

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	fmt.Println("lock start")
	time.Sleep(3 * time.Second)
	fmt.Println("lock value", c.v["somekey"])
	c.v[key]++
	fmt.Println("lock value after", c.v["somekey"])
	c.mux.Unlock()
	fmt.Println("lock end")
}


func ReadMutex()  {
	//var mutex sync.Mutex
	//fmt.Printf("%+v\n", mutex)
	////fmt.Printf("%v\n", mutex)
	//mutex.Lock()
	//fmt.Printf("%+v\n", mutex)
	//mutex.Unlock()
	//fmt.Printf("%+v\n", mutex)

	c := SafeCounter{v: make(map[string]int)}
	go c.Inc("somekey")

	time.Sleep(1 * time.Second)
	fmt.Println("inc again")
	c.v["somekey"]++
	fmt.Println("inc again", c.v["somekey"])
	fmt.Println(c.v["somekey"])

	time.Sleep(3 * time.Second)
	fmt.Println(c.v["somekey"])
}
