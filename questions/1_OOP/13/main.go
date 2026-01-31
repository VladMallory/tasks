package main

import "fmt"

type A struct{}

func (a A) Print() {
	fmt.Println("A")
}

type B struct {
	A
}

func (b B) Print() {
	// Вызов родительского метода
	b.A.Print()
	fmt.Println("B")
}

func main() {
	da := B{}
	// Да, можно
	// Если обратится к B(da), то можно
	// будет обратится к методам A
	da.A.Print() // A

	da.Print() // A B
}
