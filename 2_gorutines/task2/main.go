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

	wg.Go(func() {
		greet("Da")
	})

	wg.Wait()
}
