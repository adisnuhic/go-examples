/*
	In this example, two goroutines are communicating through two channels, ch1 and ch2. Each goroutine tries to read from one channel and write to the other.
	However, due to the order of operations, they end up sending and receiving values back and forth without making any real progress, causing a livelock.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case val := <-ch1:
				fmt.Println("Received from ch1:", val)
				ch2 <- val
			}
		}
	}()

	go func() {
		for {
			select {
			case val := <-ch2:
				fmt.Println("Received from ch2:", val)
				ch1 <- val
			}
		}
	}()

	ch1 <- 42

	time.Sleep(time.Second * 5)
}
