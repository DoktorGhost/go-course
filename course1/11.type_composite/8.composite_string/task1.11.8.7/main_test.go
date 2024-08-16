package main

import (
	"bytes"
	"os"
	"testing"
)

func TestReplaceSymbols(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		old      rune
		new      rune
		expected string
	}{
		{
			name:     "Test#1: Replace single character",
			input:    "Hello, world!",
			old:      'o',
			new:      '0',
			expected: "Hell0, w0rld!",
		},
		{
			name:     "Test#2: Replace multiple characters",
			input:    "Goodbye, world!",
			old:      'o',
			new:      'O',
			expected: "GOOdbye, wOrld!",
		},
		{
			name:     "Test#3: Replace character not present",
			input:    "Hello, world!",
			old:      'x',
			new:      'X',
			expected: "Hello, world!",
		},
		{
			name:     "Test#4: Replace all characters",
			input:    "aaaaa",
			old:      'a',
			new:      'b',
			expected: "bbbbb",
		},
		{
			name:     "Test#5: Replace with the same character",
			input:    "Test string",
			old:      's',
			new:      's',
			expected: "Test string",
		},
		{
			name:     "Test#6: Replace characters in an empty string",
			input:    "",
			old:      'a',
			new:      'b',
			expected: "",
		},
		{
			name:     "Test#7: Replace non-ASCII characters",
			input:    "Привет, мир!",
			old:      'и',
			new:      'И',
			expected: "ПрИвет, мИр!",
		},
		{
			name:     "Test#8: Replace character at the start and end",
			input:    "o Hello o",
			old:      'o',
			new:      'O',
			expected: "O HellO O",
		},
		{
			name:     "Test#9: Replace character in a string with mixed characters",
			input:    "H3ll0, w0rld!",
			old:      '0',
			new:      'O',
			expected: "H3llO, wOrld!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReplaceSymbols(tt.input, tt.old, tt.new)
			if got != tt.expected {
				t.Errorf("ReplaceSymbols(%q, %c, %c) = %q; want %q", tt.input, tt.old, tt.new, got, tt.expected)
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
	expected := "Hell0, w0rld!\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
