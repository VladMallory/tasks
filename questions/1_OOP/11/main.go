package main

import "fmt"

type B struct {
	name string
}

type A struct {
	B
}

func main() {
	a1 := A{}
	// Да, получает
	a1.name = "asd"

	fmt.Println(a1)
}
