package main

import (
	"fmt"
)

type da struct {
	name string
}

func (d da) result() da {
	return d
}

func main() {
	// будет ошибка
	// acc := bank.Account{}
	// acc.balance = 100

	// Ответ: можно сделать метод который только читает структуру
	d1 := da{"asd"}
	da := d1.result()
	fmt.Println(da)
}
