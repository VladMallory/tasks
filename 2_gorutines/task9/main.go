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

		// Тут главная суть в for
		// Если бы его не было, select выполнил бы один раз case
		// и код пошел бы дальше.
		// А так он снова и снова сюда возрвщается пока
		// не будет выполнен return
		for {
			// Входит в select один раз
			select {
			// Тут не понимаю зачем обращаться к каналу
			// который содержит time
			case <-ticker.C:
				fmt.Println("Tick")
			case <-ch:
				// Выходит из for
				return
			}
		}
	}()

	// Остановка
	time.Sleep(1500 * time.Millisecond)
	ch <- true

	fmt.Println("Stopped")
}
