package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var workers int
	ch := make(chan int)
	wg := sync.WaitGroup{}

	fmt.Scanln(&workers)

	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(<-ch)
		}()
	}

	defer func() {
		wg.Wait()
		close(ch)
	}()

	for i := 0; i < workers; i++ {
		ch <- rand.Intn(10)
	}

}
