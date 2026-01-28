package main

import "fmt"

type Engine struct{}

func (e Engine) Start() {
	fmt.Println("ААААА")
}

type Car struct {
	// Композиция
	engine Engine
}

func (c Car) Drive() {
	c.engine.Start()
}

func main() {
	car := Car{}
	car.Drive()
}
