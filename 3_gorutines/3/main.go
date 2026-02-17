package main

import (
	"fmt"
	"sync"
	"time"
)

type da1 interface {
	Inc()
	Value() int
}

type SafeCounter struct {
	mu      sync.RWMutex
	counter int
}

type SafeCounter2 struct {
	mu      sync.Mutex
	counter int
}

func (s *SafeCounter) Inc() {
	s.mu.Lock()
	s.counter++
	s.mu.Unlock()
}

func (s *SafeCounter) Value() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.counter
}

func (s *SafeCounter2) Inc() {
	s.mu.Lock()
	s.counter++
	s.mu.Unlock()
}

func (s *SafeCounter2) Value() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.counter
}

func RW() {
	start := time.Now()

	wg := sync.WaitGroup{}

	da := &SafeCounter{}

	for range 500 {

		wg.Go(func() {
			da.Inc()
		})
	}

	wg.Wait()
	println(da.Value())

	end := time.Since(start)
	fmt.Println("RW:", end)
}

func mutex() {
	start := time.Now()
	da := &SafeCounter2{}

	wg := sync.WaitGroup{}

	for range 50 {

		wg.Go(func() {
			da.Inc()
		})
	}

	wg.Wait()

	end := time.Since(start)
	fmt.Println("mutex:", end)
}

func main() {
	// В среднем 388.376µs
	RW()

	// В среднем 66.45µs
	mutex()

	// Медленее потому что логика RWMutex более
	// накладная по сравнению с Mutex
	// Когда частые записи, лучше Mutex
	// А когда много чтений, то уже лучше RWMutex
	println()
}
