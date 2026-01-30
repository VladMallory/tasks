package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	//p1 := Person{
	//	"aaa",
	//	20,
	//}
	//personClient := NewPerson(p1.name, p1.Age)
	personClient := NewPerson("asd", 12)
	fmt.Println(personClient)
}
