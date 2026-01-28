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

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	r1 := Rectangle{1, 3}
	c1 := Circle{100}
	PrintArea(r1)
	PrintArea(c1)

	//var s Shape
	//s = Circle{111}
	//fmt.Println(s.Area())
	//
	//s = Rectangle{200, 500}
	//fmt.Println(s.Area())
}
