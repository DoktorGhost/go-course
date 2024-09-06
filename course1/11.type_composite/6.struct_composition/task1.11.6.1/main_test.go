package main

import (
	"bytes"
	"os"
	"testing"
)

func TestAddDish(t *testing.T) {
	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	dish2 := Dish{Name: "Burger", Price: 5.99}

	order.AddDish(dish1)
	if len(order.Dishes) != 1 {
		t.Errorf("Expected 1 dish, got %d", len(order.Dishes))
	}
	if order.Total != 10.99 {
		t.Errorf("Expected total 10.99, got %.2f", order.Total)
	}

	order.AddDish(dish2)
	if len(order.Dishes) != 2 {
		t.Errorf("Expected 2 dishes, got %d", len(order.Dishes))
	}
	if order.Total != 16.98 {
		t.Errorf("Expected total 16.98, got %.2f", order.Total)
	}
}

func TestRemoveDish(t *testing.T) {
	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	dish2 := Dish{Name: "Burger", Price: 5.99}

	order.AddDish(dish1)
	order.AddDish(dish2)

	order.RemoveDish(dish1)
	if len(order.Dishes) != 1 {
		t.Errorf("Expected 1 dish, got %d", len(order.Dishes))
	}
	if order.Total != 5.99 {
		t.Errorf("Expected total 5.99, got %.2f", order.Total)
	}

	order.RemoveDish(dish2)
	if len(order.Dishes) != 0 {
		t.Errorf("Expected 0 dishes, got %d", len(order.Dishes))
	}
	if order.Total != 0 {
		t.Errorf("Expected total 0, got %.2f", order.Total)
	}
}

func TestCalculateTotal(t *testing.T) {
	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	dish2 := Dish{Name: "Burger", Price: 5.99}

	order.AddDish(dish1)
	order.AddDish(dish2)

	order.CalculateTotal()
	if order.Total != 16.98 {
		t.Errorf("Expected total 16.98, got %.2f", order.Total)
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
	expected := "Total: 16.98\nTotal: 5.99\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
