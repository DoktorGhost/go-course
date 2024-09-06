package main

import (
	"bytes"
	"os"
	"testing"
)

func TestGetDataString(t *testing.T) {
	tests := []struct {
		input    *bytes.Buffer
		expected string
	}{
		{
			input:    bytes.NewBufferString("Hello, World!"),
			expected: "Hello, World!",
		},
		{
			input:    bytes.NewBufferString("Go is awesome!"),
			expected: "Go is awesome!",
		},
		{
			input:    bytes.NewBufferString(""),
			expected: "",
		},
		{
			input:    bytes.NewBufferString("Multiple\nlines\nof\ntext"),
			expected: "Multiple\nlines\nof\ntext",
		},
	}

	for _, test := range tests {
		result := getDataString(test.input)
		if result != test.expected {
			t.Errorf("Expected %q, but got %q", test.expected, result)
		}
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
	expected := ""

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
