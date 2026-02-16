package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type OnceFlag struct {
	mu   sync.Mutex
	done bool
}

func (o *OnceFlag) Do(f func()) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if o.done {
		return
	}

	o.done = true
	f()
}

// Мега быстро печатает
// не принимает ...any.
func writeString(s string) {
	os.Stdout.Write([]byte(s))
	os.Stdout.Write([]byte{'\n'})
}

func writeInt(n int) {
	var buf [32]byte
	// Конвиртируем ASCII цифры (хранятся в байтах)
	b := strconv.AppendInt(buf[:0], int64(n), 10)
	b = append(b, '\n')
	os.Stdout.Write(b)
}

func printOnce(v any) {
	fmt.Fprintln(os.Stdout, v)
}

func worker(wg *sync.WaitGroup, da *OnceFlag) {
	defer wg.Done()

	da.Do(func() {
		writeString("1")
	})
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	var da OnceFlag
	var da1 int
	writeInt(da1)

	for range 5 {
		wg.Add(1)
		go worker(&wg, &da)
	}

	wg.Wait()
	printOnce(time.Since(start).String())
}
