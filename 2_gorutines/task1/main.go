package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("Main started")

	wg.Go(func() {
		fmt.Println("Goroutine finished")
	})

	wg.Wait()

	fmt.Println("Main finished")
}
