package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	count atomic.Int64
}

func (c *counter) Inc() {
	c.count.Add(1)
}

func (c *counter) Show() {
	fmt.Println(c.count.Load())
}

func main() {
	counter := counter{count: atomic.Int64{}}
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	counter.Show()
}
