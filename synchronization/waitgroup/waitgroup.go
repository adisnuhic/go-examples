package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Go routine 1")
		defer wg.Done()
	}()

	go func() {
		fmt.Println("Go routine 2")
		defer wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done...")
}
