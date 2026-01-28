package main

import "fmt"

type Animal struct{}

func (a Animal) Speak() {
	fmt.Println("Что-то было")
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Собака")
}

func (c Cat) Speak() {
	fmt.Println("Кот")
}

type Cat struct{}

func main() {
	c1 := Cat{}
	c1.Speak()

	d1 := Dog{}
	d1.Speak()
}
