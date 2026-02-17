package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {

	var mu sync.Mutex

	var wg sync.WaitGroup

	for range 100 {

		wg.Go(func() {

			// Когда идет запись нескольких го рутин на одну
			// переменную, это может вызвать ситуация когда
			// запишет сразу несколько, по этому мы должны
			// делать Mutex, чтобы читать могл
			mu.Lock()
			counter++
			mu.Unlock()
		})
	}

	wg.Wait()
	fmt.Println(counter)
}
