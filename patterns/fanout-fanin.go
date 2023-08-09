package main

import (
	"fmt"
	"sync"
)

func main() {
	tasks := make(chan string, 10)
	results := make(chan string, 10)
	var wg sync.WaitGroup

	// Create worker pool
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Send tasks to workers
	for i := 1; i <= 10; i++ {
		tasks <- fmt.Sprintf("Task %d", i)
	}
	close(tasks)

	// Wait for all workers to finish
	wg.Wait()
	close(results)

	// Collect results
	for result := range results {
		fmt.Println(result)
	}
}

func worker(id int, tasks <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		result := fmt.Sprintf("Worker %d processed %s", id, task)
		results <- result
	}
}
