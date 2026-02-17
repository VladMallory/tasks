package test

import (
	"crypto/sha256"
	"runtime"
	"strconv"
	"sync"
	printDa "task/4_gorutines/1/print"
	"time"
)

const (
	N      = 5_000_000
	rounds = 5 // усиление нагрузки
)

func heavy(i int) [32]byte {
	data := []byte(strconv.Itoa(i))
	hash := sha256.Sum256(data)

	// усиливаем CPU-нагрузку
	for range rounds {
		hash = sha256.Sum256(hash[:])
	}

	// дополнительная математика
	var x uint64
	for k := range 500 {
		x += uint64(hash[k%32]) * uint64(k+1)
	}

	return hash
}

func V1(enable bool) {
	if !enable {
		return
	}

	start := time.Now()

	for i := range N {
		_ = heavy(i)
	}

	printDa.WriteString("single: " + time.Since(start).String())
}

func V2(enable bool) {
	if !enable {
		return
	}

	start := time.Now()

	workers := runtime.NumCPU()
	jobs := make(chan int, 10_000)
	wg := sync.WaitGroup{}

	for range workers {

		wg.Go(func() {
			for i := range jobs {
				_ = heavy(i)
			}
		})
	}

	for i := range N {
		jobs <- i
	}

	close(jobs)
	wg.Wait()

	printDa.WriteString("parallel: " + time.Since(start).String())
}
