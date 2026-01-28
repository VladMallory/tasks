package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	go func() {
		// Записал в канал сначала это
		ch <- "first"
		// Потом это
		// То есть там как будто бы очередь
		ch <- "second"
	}()

	// Теперь я могу считать их по очереди
	// Сначала 1, потом 2 записаное значение
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
