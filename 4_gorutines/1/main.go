package main

import (
	"runtime"
	"sync"
	"task/4_gorutines/1/cache"
	printDa "task/4_gorutines/1/print"
	"task/4_gorutines/1/test"
	"time"
)

func main() {
	// Не выгодно печатать так много раз с горутинами
	v1(false) // 21 секунда
	v2(false) // 16 секунд

	// Чисто ради теста
	test.V1(false) // 4.6 s
	test.V2(false) // 3.1 s
}

func v2(enable bool) {
	if !enable {
		return
	}

	start := time.Now()
	cache := cache.NewCache()

	// Пишем
	cache.Set("da", "уго")

	for range 5_000_000 {
		if v, ok := cache.Get("da"); ok {
			printDa.WriteString(v)
		}
	}

	printDa.WriteString(time.Since(start).String())
}

func v1(enable bool) {
	if !enable {
		return
	}

	workers := runtime.NumCPU()
	jobs := make(chan struct{}, 10_000)
	start := time.Now()
	cache := cache.NewCache()
	wg := sync.WaitGroup{}

	// Пишем
	cache.Set("da", "уго")

	for range workers {
		go func() {
			for range jobs {
				if v, ok := cache.Get("da"); ok {
					printDa.WriteString(v)
				}
			}
		}()
	}

	for range 5_000_000 {
		jobs <- struct{}{}
	}

	close(jobs)

	wg.Wait()
	printDa.WriteString(time.Since(start).String())
}
