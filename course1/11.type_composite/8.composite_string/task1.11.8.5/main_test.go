package main

import (
	"bytes"
	"os"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Reverse standard string",
			input:    "Hello, World",
			expected: "dlroW ,olleH",
		},
		{
			name:     "Reverse numeric string",
			input:    "12345",
			expected: "54321",
		},
		{
			name:     "Reverse single character",
			input:    "1",
			expected: "1",
		},
		{
			name:     "Reverse empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Reverse single space",
			input:    " ",
			expected: " ",
		},
		{
			name:     "Reverse string with special characters",
			input:    "!@#$%^&*()",
			expected: ")(*&^%$#@!",
		},
		{
			name:     "Reverse string with unicode characters",
			input:    "Привет",
			expected: "тевирП",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseString(tt.input)
			if got != tt.expected {
				t.Errorf("ReverseString(%q) = %q; want %q", tt.input, got, tt.expected)
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
	expected := "dlroW ,olleH\n54321\n1\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
