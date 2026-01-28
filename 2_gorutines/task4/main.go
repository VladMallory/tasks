package main

import "fmt"

func main() {
	go func() {
		fmt.Println("Goroutine 1")
	}()

	go func() {
		fmt.Println("gofmt -w .Goroutine 2")
	}()

	go func() {
		fmt.Println("Goroutine 3")
	}()
}
