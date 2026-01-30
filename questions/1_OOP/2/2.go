package main

type Person struct {
	Name string
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func getPerson() *Person {
	return &Person{Name: "Bob"}
}

func main() {
	getPerson().SetName("Alice")

	// p := getPerson()
	// p.SetName("Alice")
	// fmt.Println(p.name)
}
