package bank

import "fmt"

type Account struct {
	balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Deposit amount must be greater than 0")
	}
	a.balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Withdraw amount must be greater than 0")
	}
	if a.balance < amount {
		return fmt.Errorf("Insufficient funds")
	}
	a.balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}