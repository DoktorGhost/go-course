package main

import (
	"errors"
	"fmt"
)

type Accounter interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}
type CurrentAccount struct {
	balance float64
}

func (a *CurrentAccount) Deposit(amount float64) error {
	if amount > 0 {
		a.balance += amount
		return nil
	}
	return errors.New("отрицательное число")
}
func (a *CurrentAccount) Withdraw(amount float64) error {
	if amount > a.balance {
		return errors.New("не достаточно средств")
	}
	if amount < 0 {
		return errors.New("отрицательное число")
	}
	a.balance -= amount
	return nil
}
func (a *CurrentAccount) Balance() float64 {
	return a.balance
}

type SavingsAccount struct {
	balance float64
}

func (a *SavingsAccount) Deposit(amount float64) error {
	if amount > 0 {
		a.balance += amount
		return nil
	}
	return errors.New("отрицательное число")
}
func (a *SavingsAccount) Withdraw(amount float64) error {
	if a.balance < 500 {
		return errors.New("баланс менее 500")
	}
	if amount < 0 {
		return errors.New("отрицательное число")
	}
	a.balance -= amount
	return nil
}
func (a *SavingsAccount) Balance() float64 {
	return a.balance
}

func ProcessAccount(a Accounter) {
	a.Deposit(500)
	a.Withdraw(200)
	fmt.Printf("Balance: %.2f\n", a.Balance())
}
func main() {
	c := &CurrentAccount{}
	s := &SavingsAccount{}
	ProcessAccount(c)
	ProcessAccount(s)
}
