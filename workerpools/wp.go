package main

import "fmt"

func main() {
	tasks := []int{2, 3, 4, 7, 10, 12, 40}
	workers := 3
	jobs := make(chan int, len(tasks))    //donde se van a leer todos los jobs
	results := make(chan int, len(tasks)) // donde se transmitiran los resultados
	for i := 0; i < workers; i++ {
		go worker(i, jobs, results)
	}
	for val := range tasks {
		jobs <- val
	}
	close(jobs)
	for r := 0; r < len(tasks); r++ {
		<-results //leemos los resultados
	}
}

func worker(id int, jobs <-chan int, resutls chan<- int) {
	for job := range jobs {
		fmt.Printf("worker with id %d started with %d \n", id, job)
		fib := fibonacci(job)
		fmt.Printf("worker with id %d , job %d and fib %d \n", id, job, fib)
		resutls <- fib
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
