package main

import (
	"fmt"
	"sync"
)

func foo(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; ; i++ {
		if i < 10 {
			fmt.Println("foo работает")
		} else {
			fmt.Println("foo закончила свою работу")
			return
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go foo(wg)
	wg.Wait()
}
