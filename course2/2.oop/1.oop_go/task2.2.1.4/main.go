package main

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64) error
}

type CreditCard struct {
	balance float64
}

func (cc *CreditCard) Pay(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("недопустимая сумма платежа")
	}
	if amount > cc.balance {
		return fmt.Errorf("недостаточный баланс")
	}
	cc.balance = cc.balance - amount
	fmt.Printf("Оплачено %v с помощью кредитной карты\n", amount)
	return nil
}

type Bitcoin struct {
	balance float64
}

func (b *Bitcoin) Pay(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("недопустимая сумма платежа")
	}
	if amount > b.balance {
		return fmt.Errorf("недостаточный баланс")
	}
	b.balance = b.balance - amount
	fmt.Printf("Оплачено %v с помощью кредитной карты\n", amount)
	return nil
}

func ProcessPayment(p PaymentMethod, amount float64) {
	err := p.Pay(amount)
	if err != nil {
		fmt.Println("Не удалось обработать платеж:", err)
	}
}
func main() {
	cc := &CreditCard{balance: 500.00}
	btc := &Bitcoin{balance: 2.00}
	ProcessPayment(cc, 200.00)
	ProcessPayment(btc, 1.00)
	ProcessPayment(cc, 600)
	ProcessPayment(btc, 3)
	ProcessPayment(cc, -10)
	ProcessPayment(btc, -1.00)
}
