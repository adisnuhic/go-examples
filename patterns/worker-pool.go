package main

import "fmt"

func main() {
	chJobs := make(chan int, 50)
	chResults := make(chan int, 50)

	// each go routine will pull jobs from chJobs
	go worker(chJobs, chResults)
	go worker(chJobs, chResults)
	go worker(chJobs, chResults)

	for i := 0; i < 50; i++ {
		chJobs <- i
	}

	for i := 0; i < 50; i++ {
		fmt.Println(<-chResults)
	}

}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}
