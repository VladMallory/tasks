package main

import "fmt"

// Делаю свой тип основаный на int.
type da int

// Daa Напрямую вписать int в (d int) нельзя.
// GO запрещает использовать встроенные
// типы(int, string, bool, float), но если.
// Сделать свой тип, то разрешит.
// А вот в функциях можно делать. func (i int).
func (d da) Daa(num int) int {
	return num
}

func main() {
	var x da = 10

	fmt.Println(x.Daa(50))
}
