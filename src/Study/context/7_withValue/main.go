package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type UserID string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	defer wg.Done()

	key := UserID("user_id")
	traceCode, ok := ctx.Value(key).(int64)
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code: %d\n", traceCode)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, UserID("user_id"), int64(123456))
	wg.Add(1)
	go worker(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
