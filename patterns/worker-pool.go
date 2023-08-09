package main

import "fmt"

func main() {
	tasks := make(chan string, 10)
	results := make(chan string, 10)

	// Create worker pool
	for i := 1; i <= 3; i++ {
		go worker(i, tasks, results)
	}

	// Send tasks to workers
	for i := 1; i <= 10; i++ {
		tasks <- fmt.Sprintf("Task %d", i)
	}

	close(tasks)

	// Collect results
	for i := 1; i <= 10; i++ {
		result := <-results
		fmt.Println(result)
	}
}

func worker(id int, tasks <-chan string, results chan<- string) {
	for task := range tasks {
		result := fmt.Sprintf("Worker %d processed %s", id, task)
		results <- result
	}
}
