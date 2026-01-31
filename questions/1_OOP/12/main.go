package main

import "fmt"

type A struct{}

func (a A) Print() {
	fmt.Println("A")
}

type B struct{}

func (b B) Print() {
	fmt.Println("B")
}

type C struct {
	A
	B
}

func main() {
	// Будет ошибка компилятора
	// Ambiguous reference 'Print'
	// -Неоднозначная ссылка `Print`-
	c := C{}
	// c.Print()
	c.B.Print()
}
