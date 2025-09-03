package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func foo(ch chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("функция foo отработала свое!\n")
			return
		case <-ch:
			fmt.Printf("функция foo получила значение %d\n", <-ch)
			time.Sleep(1 * time.Second)
		default:
			fmt.Printf("функция foo не получила значение!\n")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2*time.Second))
	defer cancel()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan int, 1)

	go foo(ch, wg, ctx)

	defer func() {
		wg.Wait()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- rand.Intn(10 + 1)
		}
	}

}
