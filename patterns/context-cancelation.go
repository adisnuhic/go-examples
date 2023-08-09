package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, "Worker 1")
	go worker(ctx, "Worker 2")

	time.Sleep(time.Second * 2)
	cancel()

	time.Sleep(time.Second)
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "cancelled")
			return
		default:
			fmt.Println(name, "working")
			time.Sleep(time.Second)
		}
	}
}
