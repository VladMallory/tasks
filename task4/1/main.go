package main

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct {
	Animal
}

type Cat struct {
	Animal
}

func (d *Dog) Speak() {
	fmt.Printf("Woof\n")
}

func (c *Cat) Speak() {
	fmt.Printf("мяу\n")
}

func main() {
	d1 := Dog{}
	d1.Speak()

	c1 := Cat{}
	c1.Speak()
}
