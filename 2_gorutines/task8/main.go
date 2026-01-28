package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		for i := range 5 {
			time.Sleep(200 * time.Millisecond)
			fmt.Println(i)
		}
	}()

	wg.Wait()
}
