package main

import (
	"bytes"
	"os"
	"testing"
)

func TestConcatStrings(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "Concatenate with spaces",
			input:    []string{"Hello", " ", "World"},
			expected: "Hello World",
		},
		{
			name:     "Concatenate with empty strings",
			input:    []string{"", "Hello", "", "World", ""},
			expected: "HelloWorld",
		},
		{
			name:     "Concatenate single string",
			input:    []string{"Hello"},
			expected: "Hello",
		},
		{
			name:     "Concatenate multiple strings",
			input:    []string{"Go", " ", "is", " ", "awesome"},
			expected: "Go is awesome",
		},
		{
			name:     "Concatenate empty input",
			input:    []string{},
			expected: "",
		},
		{
			name:     "Concatenate with special characters",
			input:    []string{"!@#$", "%^&*", "()"},
			expected: "!@#$%^&*()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := concatStrings(tt.input...)
			if got != tt.expected {
				t.Errorf("concatStrings(%v) = %v; want %v", tt.input, got, tt.expected)
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
	expected := "Hello World\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
