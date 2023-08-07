/*
	In this scenario, two goroutines, mu1 and mu2, are attempting to gain locks on the shared resource mu2 in the opposite order.
	The first goroutine obtains a lock on mu1, and then attempts to obtain a lock on mu2, while the second goroutine does the opposite.
	Since neither goroutine may proceed until the other releases the lock it requires, the situation is deadlocked.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu1, mu2 sync.Mutex

	wg.Add(2)
	go func() {
		defer wg.Done()
		mu1.Lock()
		fmt.Println("goroutine 1 acquired lock 1")
		mu2.Lock()
		fmt.Println("goroutine 1 acquired lock 2")
	}()
	go func() {
		defer wg.Done()
		mu2.Lock()
		fmt.Println("goroutine 2 acquired lock 2")
		mu1.Lock()
		fmt.Println("goroutine 2 acquired lock 1")
	}()
	wg.Wait()
}
