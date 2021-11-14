package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

// context内部封装了channel
func worker(ctx context.Context) {
	defer wg.Done()
	go worker2(ctx)
LABEL:
	for {
		fmt.Println("worker...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LABEL
		default:
			//...
		}
	}
}

func worker2(ctx context.Context) {
LABEL:
	for {
		fmt.Println("worker2...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LABEL
		default:
			//...
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over...")
}
