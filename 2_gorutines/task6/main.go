package main

import "fmt"

func main() {
	done := make(chan struct{})

	go func() {
		fmt.Println("Работа завершена")
		// Закрываю канал
		close(done)
	}()

	// Читаю из канала
	// Структуры можно читать сколько угодно
	fmt.Println(<-done)
	fmt.Println(<-done)
	fmt.Println(<-done)
	fmt.Println(<-done)
	fmt.Println(<-done)
	fmt.Println(<-done)

	fmt.Println("Конец")
}
