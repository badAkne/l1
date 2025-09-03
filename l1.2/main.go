package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wg.Add(len(nums))

	for i := 0; i < len(nums); i++ {
		go func(k int) {
			defer wg.Done()
			fmt.Println(k * k)
		}(i)
	}

	wg.Wait()

}
