package main

import (
	"fmt"
	"os"
	"sync"
)

func printOnce(msg any) {
	fmt.Fprintln(os.Stdout, msg)
}

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
	printOnce(sum)
}
