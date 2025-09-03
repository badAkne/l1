package main

import (
	"fmt"
	"sync"
	"time"
)

func foo(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("foo начала работу!")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	time.Sleep(2 * time.Second)
	panic("panic")
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go foo(wg)
	wg.Wait()
}
