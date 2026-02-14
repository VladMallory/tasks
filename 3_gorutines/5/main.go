package main

import (
	"fmt"
	"sync"
	"time"
)

func Marge(ch1, ch2 <-chan int) <-chan int {
	resultChan := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for v := range ch1 {
			resultChan <- v
		}
	}()

	go func() {
		defer wg.Done()
		for v := range ch2 {
			resultChan <- v
		}
	}()

	// Закрываем канал чтобы range
	// безопасно читал канал
	go func() {
		time.Sleep(time.Second * 2)
		wg.Wait()
		time.Sleep(time.Second * 2)
		close(resultChan)
		time.Sleep(time.Second * 2)
	}()

	return resultChan
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
		ch1 <- 5
	}()

	go func() {
		defer close(ch2)
		ch2 <- 0
		ch2 <- 5
		ch2 <- 10
	}()

	result := Marge(ch1, ch2)
	for v := range result {
		fmt.Println(v)
	}
}
