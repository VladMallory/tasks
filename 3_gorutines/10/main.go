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

func worker(wg *sync.WaitGroup, da *OnceFlag) {
	defer wg.Done()

	da.Do(func() {
		printOnce("1")
	})
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	var da OnceFlag

	for range 5 {
		wg.Add(1)
		go worker(&wg, &da)
	}

	wg.Wait()
	printOnce(time.Since(start).String())
}
