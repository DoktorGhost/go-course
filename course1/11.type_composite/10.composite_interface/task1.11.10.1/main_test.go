package main

import (
	"bytes"
	"os"
	"testing"
)

func TestGetType(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{42, "int"},
		{3.14, "float"},
		{"hello", "string"},
		{[]int{1, 2, 3}, "[]int"},
		{[]float64{1.1, 2.2, 3.3}, "[]float64"},
		{[]string{"a", "b", "c"}, "[]string"},
		{nil, "Пустой интерфейс"},
		{struct{}{}, "Тип неопределен"}, // структура без полей как пример неопределенного типа
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := getType(tt.input)
			if result != tt.expected {
				t.Errorf("getType(%v) = %q; want %q", tt.input, result, tt.expected)
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
	expected := "int\nstring\n[]int\nПустой интерфейс\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
