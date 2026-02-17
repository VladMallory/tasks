package main

import (
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan struct{}, 3)

	for i := range 10 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			ch <- struct{}{}

			time.Sleep(time.Millisecond * 300)

			println("конец", id)
			<-ch
		}(i)
	}

	wg.Wait()
}
