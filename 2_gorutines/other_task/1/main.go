package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 1; i <= 100; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()

		sum := 0
		for i := 1; i <= 100; i++ {
			sum = <-ch
		}

		fmt.Println(sum)
	}()

	wg.Wait()
}
