package main

import "fmt"

type Counter struct {
	value int
}

type CounterInterface interface {
	GetValue() int
	Increment()
	SetValue(newValue int)
}

// GetValue возвращает текущее значение.
func (c *Counter) GetValue() int {
	return c.value
}

func (c *Counter) Increment() {
	c.value++
}

// SetValue устанавливает новое значение.
func (c *Counter) SetValue(newValue int) {
	c.value = newValue
}

func main() {
	// Так работает, мы можем менять взаимодействовать
	// через методы
	var ci CounterInterface = &Counter{10}

	fmt.Println("Начало", ci.GetValue())

	// Изменяем состояние
	ci.Increment()
	fmt.Println("После Increment", ci.GetValue())

	// Новое значение
	ci.SetValue(20)

	fmt.Println("После SetValue", ci.GetValue())
	// А вот так нельзя делать. Напрямую обращаться
	// Будем ошибка компилятора
	// ci.value = 30 ошибка комплятора, через интерфейс нельзя менять

	fmt.Println()
}
