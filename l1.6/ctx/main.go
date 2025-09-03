package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func foo(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("foo закончила работу")
			return
		default:
			fmt.Println("Функция foo работает")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3*time.Second))
	defer cancel()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go foo(ctx, wg)
	wg.Wait()
}
