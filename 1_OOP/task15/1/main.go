package main

import "fmt"

type List struct{}

// Len расчитывает длину строки.
func (l *List) Len(msg string) int {
	if l == nil {
		return 0
	}

	return len(msg)
}

func main() {
	l1 := List{}
	fmt.Println(l1.Len("asd"))
	fmt.Println(l1.Len(""))
}
