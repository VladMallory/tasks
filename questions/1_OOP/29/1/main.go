package main

import "fmt"

type Counter struct {
	value int
}

// Метод с приемником-указателем.
func (c *Counter) Increment() {
	if c == nil {
		fmt.Println("Приемник nil, ничего не делаем")

		return
	}
	c.value++
}

func (c *Counter) Value() int {
	if c == nil {
		return 0 // Возвращаем дефолтное значение для nil-объекта
	}

	return c.value
}

func main() {
	var c *Counter = nil // c равен nil

	// Вызов метода на nil безопасно
	c.Increment()
	fmt.Println("Value:", c.Value()) // Вывод: Value: 0
}
