package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Greet выводит имя пользователя из структуры
func (p *Person) Greet() {
	fmt.Printf("Hello, I'm %s\n", p.Name)
}

func main() {
	// Создаем ЧЕЛОВЕКА
	p1 := Person{
		Name: "asd",
		Age:  20,
	}

	p1.Greet()
}
