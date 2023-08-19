package banking

import "math"

type Account struct {
	ID      int
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount float64) bool {
	a.Balance -= amount
	a.Balance = math.Abs(a.Balance)
	return true
}
