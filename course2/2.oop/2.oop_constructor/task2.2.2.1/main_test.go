package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"
)

// Функция для сравнения двух объектов Order
func ordersEqual(o1, o2 Order) bool {
	return o1.ID == o2.ID &&
		o1.CustomerID == o2.CustomerID &&
		equalStringSlices(o1.Items, o2.Items) &&
		o1.OrderDate.Equal(o2.OrderDate)
}

// Вспомогательная функция для сравнения срезов строк
func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Тестируем создание заказа с установленным ID и опциями
func TestNewOrder(t *testing.T) {
	// Установим текущую дату и время для использования в тесте
	now := time.Now()

	tests := []struct {
		id          int
		options     []OrderOption
		expected    Order
		description string
	}{
		{
			id: 1,
			options: []OrderOption{
				WithCustomerID("123"),
				WithItems([]string{"item1", "item2"}),
				WithOrderDate(now),
			},
			expected: Order{
				ID:         1,
				CustomerID: "123",
				Items:      []string{"item1", "item2"},
				OrderDate:  now,
			},
			description: "Order with ID, customer ID, items, and order date",
		},
		{
			id:          2,
			options:     []OrderOption{},
			expected:    Order{ID: 2},
			description: "Order with only ID set",
		},
		{
			id: 3,
			options: []OrderOption{
				WithCustomerID("456"),
			},
			expected: Order{
				ID:         3,
				CustomerID: "456",
			},
			description: "Order with ID and customer ID",
		},
		{
			id: 4,
			options: []OrderOption{
				WithItems([]string{"item3"}),
			},
			expected: Order{
				ID:    4,
				Items: []string{"item3"},
			},
			description: "Order with ID and items",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			order := NewOrder(test.id, test.options...)
			if !ordersEqual(*order, test.expected) {
				t.Errorf("NewOrder(%d, %v) = %+v; want %+v", test.id, test.options, order, test.expected)
			}
		})
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
	expected := "Order: &{ID:1 CustomerID:123 Items:[item1 item2] OrderDate:"

	if strings.Contains(stdout.String(), expected) == false {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
