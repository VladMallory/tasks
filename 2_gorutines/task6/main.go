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
	<-done

	fmt.Println("Конец")
}
