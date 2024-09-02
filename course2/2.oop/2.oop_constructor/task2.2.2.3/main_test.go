package main

import (
	"bytes"
	"os"
	"testing"
)

func TestSavingsAccount(t *testing.T) {
	// Создаем новый SavingsAccount
	account := &SavingsAccount{}

	// Тестируем депозиты
	account.Deposit(1000)
	if got := account.Balance(); got != 1000 {
		t.Errorf("Deposit failed: got %v, want %v", got, 1000)
	}

	err := account.Withdraw(1500)
	if err == nil {
		t.Errorf("Withdraw failed: %v", err)
	}

	// Тестируем снятие средств (успешное)
	err = account.Withdraw(500)
	if err != nil {
		t.Errorf("Withdraw failed: %v", err)
	}
	if got := account.Balance(); got != 500 {
		t.Errorf("Balance after withdrawal failed: got %v, want %v", got, 500)
	}

	// Тестируем снятие средств (ошибка из-за баланса менее 1000)
	err = account.Withdraw(600)
	if err == nil {
		t.Error("Expected error when withdrawing with balance less than 1000, got nil")
	}

	// Тестируем депозит и снятие в отрицательном балансе
	account.Deposit(-100)
	if got := account.Balance(); got != 500 {
		t.Errorf("Deposit with negative amount failed: got %v, want %v", got, 500)
	}

	err = account.Withdraw(-600)
	if err == nil {
		t.Error("Expected error when withdrawing with balance less than 1000, got nil")
	}
}

func TestCheckingAccount(t *testing.T) {
	// Создаем новый CheckingAccount
	account := &CheckingAccount{}

	// Тестируем депозиты
	account.Deposit(1000)
	if got := account.Balance(); got != 1000 {
		t.Errorf("Deposit failed: got %v, want %v", got, 1000)
	}

	// Тестируем снятие средств (успешное)
	err := account.Withdraw(500)
	if err != nil {
		t.Errorf("Withdraw failed: %v", err)
	}
	if got := account.Balance(); got != 500 {
		t.Errorf("Balance after withdrawal failed: got %v, want %v", got, 500)
	}

	// Тестируем снятие средств (ошибка из-за недостатка средств)
	err = account.Withdraw(600)
	if err == nil {
		t.Error("Expected error when withdrawing more than balance, got nil")
	}

	err = account.Withdraw(-600)
	if err == nil {
		t.Error("Expected error when withdrawing with balance less than 1000, got nil")
	}
}

func TestNewCustomer(t *testing.T) {
	// Создаем SavingsAccount
	account := &SavingsAccount{}
	account.Deposit(2000)

	// Создаем нового клиента
	customer := NewCustomer(1, WithName("Alice"), WithAccount(account))

	// Проверяем, что клиент создан правильно
	if customer.ID != 1 {
		t.Errorf("Customer ID failed: got %v, want %v", customer.ID, 1)
	}
	if customer.Name != "Alice" {
		t.Errorf("Customer Name failed: got %v, want %v", customer.Name, "Alice")
	}
	if got := customer.Account.Balance(); got != 2000 {
		t.Errorf("Customer Account Balance failed: got %v, want %v", got, 2000)
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	var stdout bytes.Buffer
	stdout.ReadFrom(r)
	expected := "отрицательное число\nCustomer: John Doe, Balance: 1000\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
