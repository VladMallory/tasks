package main

import (
	"fmt"
	"sync"
)

func greet(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		greet("Da")
	}()

	wg.Wait()
}
