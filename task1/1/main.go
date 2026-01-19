package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type bankAccount struct {
	// Я слышал что баланс лучше хранить в int
	// в float32/64 какие-то проблемы
	balance int
}

// Deposit пополняет баланс, нужно указать сумму в amount
func (b *bankAccount) Deposit(amount int) {
	// Пополняем баланс
	b.balance = b.balance + amount
}

func (b *bankAccount) GetBalance() int {
	return b.balance
}

// Greet выводит имя пользователя из структуры
func (p *Person) Greet() {
	fmt.Printf("Hello, I'm %s\n", p.Name)
}

// Birthday обращается к структуре и меняет значение
// Ссылаемся на тот же участок памяти,
// прибавляем +1 и записываем в тот же участок памяти 21
// (p *Person) - создается локальная переменная p которая ссылается
// в плане памяти на Person, то есть на один и тот же
// участок кода
func (p *Person) Birthday() {
	p.Age++
	//p.Age = p.Age + 1
}

// BirthdayValue делает копию структуры в p и меняет в копии
// значение. Оригнал не меняется
// p так же локальная переменная. Но содержит теперь не содержит ссылку
// на оригнал, а содержит ссылку на копию структуры Person
func (p Person) BirthdayValue() {
	p.Age++
}

func main() {
	// Создаем ЧЕЛОВЕКА
	p1 := Person{
		Name: "asd",
		Age:  20,
	}

	p1.Birthday()
	p1.Greet()
	p1.BirthdayValue()

	b1 := bankAccount{
		balance: 200,
	}

	// Пополняем баланс и узнаем баланс
	b1.Deposit(200)
	resultBalance := b1.GetBalance()
	fmt.Println(resultBalance)

	b1.Deposit(200)

	resultBalance2 := b1.GetBalance()
	fmt.Println(resultBalance2)
}
