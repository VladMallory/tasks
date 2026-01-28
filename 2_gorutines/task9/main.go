package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	go func() {
		ticker := time.NewTicker(300 * time.Millisecond)
		// Останавливает ticker
		// Так не будет утечки
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Println("Tick")
			case <-ch:
				return
			}
		}
	}()
}
