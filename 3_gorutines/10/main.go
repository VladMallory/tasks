package main

import (
	"fmt"
	"os"
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

func printOnce(msg string) {
	fmt.Fprintln(os.Stdout, msg)
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	var da OnceFlag

	wg.Add(5)

	for range 5 {
		go func() {
			defer wg.Done()

			da.Do(func() {
				printOnce("1")
			})
		}()
	}

	wg.Wait()
	printOnce(time.Since(start).String())
}
