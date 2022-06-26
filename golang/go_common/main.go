package main

import (
	"fmt"
	"sync"
)

func main1() {
	var (
		wg       = sync.WaitGroup{}
		errChan  = make(chan error, 1)
		finished = make(chan bool, 1)
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		errChan <- fmt.Errorf("test")
		fmt.Println("errChan <-")
	}()
	go func() {
		wg.Wait()
		finished <- false
		close(finished)
		fmt.Println("close(finished)")
	}()
	// for {
	// 	select {
	// 	case v, isClosed := <-finished:
	// 		fmt.Printf("finised, %v, %v\n", v, isClosed)
	// 		return
	// 	case err := <-errChan:
	// 		fmt.Printf("<-errChan: %s\n", err)
	// 	}
	// }

	// for k := range finished {
	// 	fmt.Printf("k=%v", k)
	// 	fmt.Printf("<-errChan: %s\n", <-errChan)
	// }
	// for err := range errChan {
	// 	fmt.Printf("%v", err)
	// }
}
