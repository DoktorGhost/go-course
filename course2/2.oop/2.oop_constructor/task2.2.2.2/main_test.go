package main

import (
	"bytes"
	"os"
	"testing"
)

func TestNewOrder(t *testing.T) {
	tests := []struct {
		name     string
		id       int
		options  []UserOption
		expected User
	}{
		{
			name: "Test 1",
			id:   1,
			options: []UserOption{
				WithUsername("Fedor"),
				WithEmail("123@qq.com"),
				WithRole("admin"),
			},
			expected: User{
				ID:       1,
				Username: "Fedor",
				Email:    "123@qq.com",
				Role:     "admin",
			},
		},
		{
			name:     "Test 2",
			id:       2,
			options:  []UserOption{},
			expected: User{ID: 2},
		},
		{
			name: "Test 3",
			id:   3,
			options: []UserOption{
				WithUsername("Ivan"),
			},
			expected: User{
				ID:       3,
				Username: "Ivan",
			},
		},
		{
			id: 4,
			options: []UserOption{
				WithEmail("123@qq.com"),
			},
			expected: User{
				ID:    4,
				Email: "123@qq.com",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			order := NewUser(test.id, test.options...)
			if *order != test.expected {
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
	expected := "User: &{ID:1 Username:testuser Email:testuser@example.com Role:admin}\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
