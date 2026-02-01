package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var ch chan int = make(chan int, 5)

	var wg sync.WaitGroup

	start := time.Now()

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 1; i <= 200000; i++ {
			ch <- i
		}

		close(ch)
	}()

	go func() {
		defer wg.Done()

		for v := range ch {
			fmt.Println(v)
		}
	}()

	wg.Wait()

	fmt.Println(time.Since(start))
}
