package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Worker interface {
	Do(ctx context.Context) error
}

type FileProcessor struct{}

type NetworkFetcher struct{}

func (f FileProcessor) Do(ctx context.Context) error {
	// Проверяем была ли ошибка в context
	if ctx.Err() != nil {
		return ctx.Err()
	}

	return nil
}

func (n NetworkFetcher) Do(ctx context.Context) error {
	// Проверяем была ли ошибка в context
	if ctx.Err() != nil {
		return ctx.Err()
	}

	// Имитация сетевой работы
	time.Sleep(1 * time.Second)

	return nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		da1()
	}()

	go func() {
		defer wg.Done()
		da2()
	}()

	wg.Wait()

}

func da1() {

	// закрываю соединение
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// Воркер
	var worker Worker

	// FileProcessor
	worker = FileProcessor{}

	// Обращаюсь к интерфейсу и его
	// методу do, даю туда контекс
	// Так мы привязываем это к интерфейсу и можно делать
	// подмену реализаций метода Do()
	if err := worker.Do(ctx); err != nil {
		fmt.Println(err)
	} else {
		// Поскольку ошибка nil, будет успешно
		fmt.Println("FileProcessor: успешно")
	}
}

// Это для эксперемента. Поскольку службы между собой
// не связаны, можно сделать асинхронно
func da2() {
	ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	// закрываю соединение
	defer cancel2()

	worker := NetworkFetcher{}
	if err := worker.Do(ctx2); err != nil {
		fmt.Println(err)
	} else {
		// Поскольку ошибка nil, будет успешно
		fmt.Println("NetworkFetcher: успешно")
	}
}
