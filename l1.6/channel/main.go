package main

import (
	"fmt"
	"sync"
	"time"
)

func foo(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ch:
			fmt.Println("foo закончила работу")
			return
		default:
			fmt.Println("foo работает")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan bool)

	wg.Add(1)
	go foo(ch, wg)

	time.Sleep(3 * time.Second)
	ch <- true

	wg.Wait()
}
