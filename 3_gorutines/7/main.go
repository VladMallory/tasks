package main

import (
	"fmt"
	"os"
	"sync"
)

type SafeLogger struct {
	mu  sync.Mutex
	out *os.File
}

func (l *SafeLogger) Log(msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, err := fmt.Fprintf(l.out, msg)
	if err != nil {
		return
	}
}

func main() {
	da := &SafeLogger{out: os.Stdout}
	wg := sync.WaitGroup{}

	for i := 0; i < 500000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			da.Log(fmt.Sprintf("медленно жесть. Столько fmt. %d\n", id))
		}(i)
	}

	wg.Wait()
}
