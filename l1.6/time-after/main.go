package main

import (
	"fmt"
	"sync"
	"time"
)

func foo(wg *sync.WaitGroup, timeoutCh <-chan time.Time) {
	defer wg.Done()
	for {
		select {
		case <-timeoutCh:
			fmt.Println("функция foo отработала!")
			return
		default:
			fmt.Println("Функция foo работает")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	timeoutCh := time.After(3 * time.Second)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go foo(wg, timeoutCh)
	wg.Wait()
}
