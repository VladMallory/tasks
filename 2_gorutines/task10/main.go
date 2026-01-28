package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)

	wg.Add(2)

	go func() {
		defer wg.Done()
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	go func() {
		defer wg.Done()
		// Можно еще так
		// da1 := <-ch
		// da2 := <-ch
		// da3 := <-ch
		// fmt.Println(da1, da2, da3)

		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}()

	wg.Wait()
}
