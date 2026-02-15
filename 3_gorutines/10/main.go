package main

import (
	"fmt"
	"sync"
	"time"
)

type OnceFlag struct {
	mu   sync.Mutex
	done bool
}

func (o *OnceFlag) Do(f func()) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if o.done {
		return
	}

	o.done = true
	f()
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	var da OnceFlag

	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()

			da.Do(func() {
				fmt.Println("1")
			})
		}()
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}
