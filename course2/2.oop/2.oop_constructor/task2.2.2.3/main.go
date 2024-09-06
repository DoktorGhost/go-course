package main

import (
	"fmt"
	"sync"
)

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	Balance() float64
}

type Customer struct {
	ID      int
	Name    string
	Account Account
}

type CustomerOption func(*Customer)

func NewCustomer(id int, opts ...CustomerOption) *Customer {
	u := &Customer{ID: id}
	for _, opt := range opts {
		opt(u)
	}
	return u
}

func WithName(name string) CustomerOption {
	return func(u *Customer) {
		u.Name = name
	}
}

func WithAccount(account Account) CustomerOption {
	return func(u *Customer) {
		u.Account = account
	}
}

type SavingsAccount struct {
	balance float64
	mu      sync.Mutex
}

func (sa *SavingsAccount) Deposit(amount float64) {
	sa.mu.Lock()
	defer sa.mu.Unlock()
	if amount >= 0 {
		sa.balance += amount
	}
}

func (sa *SavingsAccount) Withdraw(amount float64) error {
	sa.mu.Lock()
	defer sa.mu.Unlock()
	if sa.balance < 1000 {
		return fmt.Errorf("баланс менее 1000")
	}
	if amount > sa.balance {
		return fmt.Errorf("на балансе не достаточно средств")
	}
	if amount < 0 {
		return fmt.Errorf("отрицательное число")
	}
	sa.balance -= amount
	return nil
}
func (sa *SavingsAccount) Balance() float64 {
	sa.mu.Lock()
	defer sa.mu.Unlock()
	return sa.balance
}

type CheckingAccount struct {
	balance float64
	mu      sync.Mutex
}

func (ca *CheckingAccount) Deposit(amount float64) {
	ca.mu.Lock()
	defer ca.mu.Unlock()
	if amount >= 0 {
		ca.balance += amount
	}
}

func (ca *CheckingAccount) Withdraw(amount float64) error {
	ca.mu.Lock()
	defer ca.mu.Unlock()
	if amount > ca.balance {
		return fmt.Errorf("на балансе не достаточно средств")
	}
	if amount < 0 {
		return fmt.Errorf("отрицательное число")
	}
	ca.balance -= amount
	return nil
}
func (ca *CheckingAccount) Balance() float64 {
	ca.mu.Lock()
	defer ca.mu.Unlock()
	return ca.balance
}

func main() {
	savings := &SavingsAccount{}
	savings.Deposit(1000)
	customer := NewCustomer(1, WithName("John Doe"), WithAccount(savings))
	err := customer.Account.Withdraw(-100)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Customer: %v, Balance: %v\n", customer.Name, customer.Account.Balance())
}
