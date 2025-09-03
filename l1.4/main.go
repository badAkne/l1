package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func foo(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("foo номер: %d заканчивает работу\n", id)
			time.Sleep(1 * time.Second)
			return
		default:
			fmt.Printf("foo номер: %d продолжает работу\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	var workers int
	fmt.Scan(&workers)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)

	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go foo(ctx, wg, i)
	}

	<-sigCh
	cancel()

	wg.Wait()
}
