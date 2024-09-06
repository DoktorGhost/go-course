package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID         int
	CustomerID string
	Items      []string
	OrderDate  time.Time
}
type OrderOption func(*Order)

func WithCustomerID(customerID string) OrderOption {
	return func(o *Order) {
		o.CustomerID = customerID
	}
}

func WithItems(arr []string) OrderOption {
	return func(o *Order) {
		o.Items = arr
	}
}

func WithOrderDate(time time.Time) OrderOption {
	return func(o *Order) {
		o.OrderDate = time
	}
}

func NewOrder(id int, opts ...OrderOption) *Order {
	o := &Order{ID: id}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func main() {
	order := NewOrder(1,
		WithCustomerID("123"),
		WithItems([]string{"item1", "item2"}),
		WithOrderDate(time.Now()))
	fmt.Printf("Order: %+v\n", order)
}
