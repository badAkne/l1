package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var mtx = &sync.Mutex{}

func Put(m map[int]int, k, v int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer mtx.Unlock()
	mtx.Lock()
	m[k] = v
}

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Put(m, rand.Intn(10), rand.Intn(10), wg)
	}

	wg.Wait()
	for k, v := range m {
		fmt.Println(k, v)
	}
}
