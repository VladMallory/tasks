package main

import (
	"context"
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
	//f1 := FileProcessor{}
	//
	//n1 := NetworkFetcher{}

}
