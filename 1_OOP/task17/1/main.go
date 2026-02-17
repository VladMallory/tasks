package main

import "fmt"

func ConvertIntsToInterface(ints []int) []any {
	result := make([]any, len(ints))

	// складывает в result все что получил из ints
	for i, v := range ints {
		result[i] = v
	}

	// и возвраем итоговый слайс. А поскольку у нас
	// []interface (или any), можно положить слайс
	return result
}

func main() {
	slice := []int{1, 2, 4, 5, 5, 100, 2}

	// []int и interface это разные типы. Мы не можем напрямую
	// их преобразовать
	da := ConvertIntsToInterface(slice)
	fmt.Println(da)
}
