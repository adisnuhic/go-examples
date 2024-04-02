// The goal of this code is to demonstrate how to gracefully shutdown go routines

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	doneCh := make(chan bool)

	// If the main crashes it will close the doneCh, and all go routines will grcefully shutdown so no go routine leaks
	defer close(doneCh)

	cows := make(chan interface{}, 100)
	pigs := make(chan interface{}, 100)

	go func() {
		for {
			select {
			case <-doneCh:
				return
			case cows <- "moo":
			}
		}
	}()

	go func() {
		for {
			select {
			case <-doneCh:
				return
			case pigs <- "oink":
			}
		}
	}()

	wg.Add(1)
	go consumeCows(doneCh, cows, &wg)

	wg.Add(1)
	go consumePigs(doneCh, pigs, &wg)

	wg.Wait()

}

func consumeCows(doneCh chan bool, cows <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-doneCh:
			return
		case cow, ok := <-cows:
			if !ok {
				return
			}
			fmt.Println(cow)
		}
	}
}

func consumePigs(doneCh chan bool, pigs <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-doneCh:
			return
		case pg, ok := <-pigs:
			if !ok {
				return
			}
			fmt.Println(pg)
		}
	}
}
