package main

import "fmt"

type IntList struct {
	Value int
	Next  *IntList
}

// Метод Sum суммирует все элементы списка.
func (l *IntList) Sum() int {
	// Самое важное происходит здесь:
	if l == nil {
		// Если список не существует (nil), считаем, что сумма равна 0.
		// Это логично: сумма пустого множества равна нулю.
		return 0
	}

	// Если список существует, складываем текущее значение и сумму остатка.
	// Здесь мы рекурсивно вызываем Sum() у Next.
	// Если Next равен nil, следующий вызов вернет 0 (базовый случай выше).
	return l.Value + l.Next.Sum()
}

func main() {
	// Создаем список: 1 -> 2 -> 3 -> nil
	list := &IntList{Value: 1, Next: &IntList{Value: 2, Next: &IntList{Value: 3}}}
	fmt.Println(list.Sum()) // Вывод: 6

	// А теперь главный фокус: список равен nil
	var emptyList *IntList = nil

	// Мы вызываем метод на nil!
	// Это безопасно, потому что внутри есть проверка.
	fmt.Println(emptyList.Sum()) // Вывод: 0
}
