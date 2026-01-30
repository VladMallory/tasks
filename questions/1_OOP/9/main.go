package main

import "fmt"

type da struct {
	name string
}

func (d da) result() da {
	return d
}

func main() {
	// Ответ: можно сделать метод который только читает структуру
	d1 := da{"asd"}
	da := d1.result()
	fmt.Println(da)
}
