package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Stats struct {
	requests int
	errors   int
	mu       sync.Mutex
}

func (s *Stats) RecordRequest() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.requests++
}

func (s *Stats) RecordError() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.errors++
}

func main() {
	// Генерирует случайные числа чтобы
	// имитировать более живое поведение
	// time.Now - текущее время
	// UnixNano - возвращает количество наносекунд с 1 января 1970 года
	// используется как seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	wg := sync.WaitGroup{}
	da := &Stats{}

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				da.RecordRequest()

				if r.Float64() < 0.1 {
					da.RecordError()
				}
			}
		}()
	}

	wg.Wait()
	fmt.Println(da.errors)
	fmt.Println(da.requests)
}
