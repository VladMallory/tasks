package main

import "fmt"

type MyInt int

func (m MyInt) String() string {
	return fmt.Sprintf("dada %d", m)
}

func main() {
	var x any = MyInt(42)
	fmt.Println(x)
}
