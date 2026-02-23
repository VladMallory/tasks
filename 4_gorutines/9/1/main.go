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

	// Пишем редко
	wg.Add(1)
	go func() {
		defer wg.Done()

		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				da.Add(time.Now())
				fmt.Println("пишем", time.Now().Format("15:04:05"))
			}
		}
	}()

	// Много читаем за итерацию
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			ticker := time.NewTicker(2 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					events := da.GetLast()
					fmt.Println("читаем", id, len(events))
				}
			}
		}(i)
	}

	// Закрываем после того как все синхронизировали
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
