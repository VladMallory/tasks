package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	da := NewTimestampStore(3)

	var wg sync.WaitGroup
	done := make(chan bool)

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for range 10 {
				da.Add(time.Now())
				time.Sleep(time.Second * 2)
			}
		}()
	}

	go func() {
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				// Вызываем GetLast для проверки
				events := da.GetLast()
				fmt.Printf(
					"Чтение: %s, событий в буфере: %d\n",
					time.Now().Format("15:04:05"),
					len(events),
				)
			}
		}
	}()

	go func() {
		wg.Wait()
		close(done)
	}()

	<-done
}

type TimestampStore struct {
	mu      sync.RWMutex
	events  []time.Time
	maxSize int
}

func NewTimestampStore(size int) *TimestampStore {
	return &TimestampStore{
		events:  make([]time.Time, 0, size),
		maxSize: size,
	}
}

func (s *TimestampStore) Add(t time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.events = append(s.events, t)

	if len(s.events) > s.maxSize {
		s.events = s.events[1:]
	}
}

func (s *TimestampStore) GetLast() []time.Time {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]time.Time, len(s.events))
	copy(result, s.events)

	return result
}
