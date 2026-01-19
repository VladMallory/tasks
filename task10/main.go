package main

import "fmt"

type da1 interface {
	Shape()
}

type Square struct {
	a, b int
}

type Da2 struct {
	a, b int
}

type Da3 struct {
	a, b int
}

func (s Square) Shape() {
	fmt.Println(s.a * s.b)
}

func (s Da2) Shape() {
	fmt.Println(s.a * s.b)
}

func (s Da3) Shape() {
	fmt.Println(s.a * s.b)
}

func check(i any) {
	// Есть ли где-то в интерфейсе структура Square
	// Достали структуру Square нашей проверкой,
	// а проверка, есть ли там структура Square
	sq, ok := i.(Square)
	// Если true
	if ok {
		fmt.Println("Да1")
		sq.Shape()
	} else {
		sq.Shape()
		fmt.Println("Что за ужас тут творится")
	}

	sq1, ok := i.(Da2)
	if ok {
		sq1.Shape()
	} else {
		fmt.Println("хммм")
	}
}

func main() {
	// Вот тут мы создаем интерфейс
	// и в нем уже лежит структура Square
	var s da1 = Da3{100, 5}
	// Мы пихаем в check интерфейс который
	// уже содержит в себе структуру
	check(s)
}
