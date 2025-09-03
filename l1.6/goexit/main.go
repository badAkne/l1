package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func foo(wg *sync.WaitGroup) {
	defer func(*sync.WaitGroup) {
		defer fmt.Println("foo закончила работу!")
		defer wg.Done()
	}(wg)
	fmt.Println("foo начала работу!")
	time.Sleep(2 * time.Second)
	runtime.Goexit()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go foo(wg)
	wg.Wait()
}
