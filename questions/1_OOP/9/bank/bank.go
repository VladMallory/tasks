package bank

type Account struct {
	balance float64
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
}
