package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func writer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var nums [10]int
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(10) + 1
	}
	for _, num := range nums {
		ch <- num
	}
	close(ch)
}

func multiply(ch1, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var res int

	for num := range ch {
		res = num * num
		ch1 <- res
	}
	close(ch1)
}

func main() {
	ch := make(chan int)
	ch1 := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go writer(ch, wg)
	go multiply(ch1, ch, wg)

	for num := range ch1 {
		fmt.Println(num)
	}
	wg.Wait()
}
