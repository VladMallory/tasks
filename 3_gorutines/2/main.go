package main

import (
	"fmt"
	"sync"
)

type da1 interface {
	Inc()
	Value() int
}

type SafeCounter struct {
	mu      sync.Mutex
	counter int
}

func (s *SafeCounter) Inc() {
	s.mu.Lock()
	s.counter++
	s.mu.Unlock()
}

func (s *SafeCounter) Value() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.counter
}

func main() {
	wg := sync.WaitGroup{}

	da := &SafeCounter{}

	for range 50 {

		wg.Go(func() {
			da.Inc()
		})
	}

	wg.Wait()
	fmt.Println(da.Value())
}
