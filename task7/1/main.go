package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	//c1 := Circle{111}
	//r1 := Rectangle{110, 10}

	var s Shape
	s = Circle{111}
	fmt.Println(s.Area())

	s = Rectangle{200, 500}
	fmt.Println(s.Area())

	//fmt.Println("c1", c1.Area())
	//fmt.Println("r1", r1.Area())
}
