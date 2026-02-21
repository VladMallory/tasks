package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println("Animal")
}

type Dog struct {
	Animal // Встраивание
	Name   string
}

func (d Dog) Speak() {
	fmt.Println("Dog")
}

func main() {
	da1 := Dog{
		Animal: Animal{Name: "da"},
		Name:   "da1",
	}

	da1.Animal.Speak()
	da1.Speak()
}
