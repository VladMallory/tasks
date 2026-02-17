package main

import (
	"sync"
	"task/4_gorutines/2/counter"
)

func main() {
	da := counter.NewAtomicCounter()
	numG := 5_000_000
	wg := sync.WaitGroup{}

	wg.Add(numG)

	for range numG {
		go func() {
			defer wg.Done()
			da.Increment()
		}()
	}

	wg.Wait()
}
