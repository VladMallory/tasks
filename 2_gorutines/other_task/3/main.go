package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println(id, job)

		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	jobs := make(chan int)

	var wg sync.WaitGroup

	const workers = 3

	wg.Add(workers)

	for i := 1; i <= workers; i++ {
		go worker(i, jobs, &wg)
	}

	// Продюсер
	for j := 1; j <= 20; j++ {
		jobs <- j
	}

	close(jobs)

	wg.Wait()
}
