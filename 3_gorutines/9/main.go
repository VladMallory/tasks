package main

import (
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	sum := 0

	wg.Add(2)

	// Генератор
	go func() {
		defer wg.Done()

		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	// Сумирует числа
	go func() {
		defer wg.Done()
		for v := range ch {
			sum += v
		}
	}()

	wg.Wait()
	println(sum)
}
