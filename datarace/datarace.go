/*
	Two Goroutines are concurrently trying to modify the counter variable by incrementing and decrementing it.
	Since there is no synchronization mechanism like a mutex or a channel to control access to the shared variable,
	a data race can occur. The order of execution of these two Goroutines is unpredictable.
	Run: go run -race datarace/datarace.go
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	var counter int

	go func() {
		counter++
	}()

	go func() {
		counter--
	}()

	// Wait for Goroutines to finish
	time.Sleep(time.Millisecond * 100)

	fmt.Println("Final counter value:", counter)
}

// fix for data race
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var counter int
// 	var mx sync.Mutex
// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	go func() {
// 		mx.Lock()
// 		counter++
// 		defer mx.Unlock()
// 		defer wg.Done()
// 	}()

// 	go func() {
// 		mx.Lock()
// 		counter--
// 		defer mx.Unlock()
// 		defer wg.Done()
// 	}()

// 	wg.Wait()

// 	fmt.Println("Final counter value:", counter)
// }
