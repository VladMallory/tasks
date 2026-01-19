package main

import (
	"fmt"
)

type Speaker interface {
	Speak()
}

func MakeSound(s Speaker) {
	s.Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Гав")
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Мяу")
}

func main() {
	animals := []Speaker{Dog{}, Cat{}}
	fmt.Println(animals)

	for _, a := range animals {
		MakeSound(a)
	}

	//d1 := Dog{}
	//d1.Speak()
	//
	//c1 := Cat{}
	//fmt.Println()
	//c1.Speak()
}
