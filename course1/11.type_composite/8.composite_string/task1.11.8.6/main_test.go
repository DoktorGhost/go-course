package main

import (
	"bytes"
	"os"
	"testing"
)

func TestCountVowels(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Test#1: Empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "Test#2: String with no vowels",
			input:    "bcd",
			expected: 0,
		},
		{
			name:     "Test#3: String with only vowels",
			input:    "aeiou",
			expected: 5,
		},
		{
			name:     "Test#4: String with mixed vowels and consonants",
			input:    "Hello, World!",
			expected: 3,
		},
		{
			name:     "Test#5: String with Cyrillic vowels",
			input:    "Привет, Мир!",
			expected: 3,
		},
		{
			name:     "Test#6: String with mixed Latin and Cyrillic vowels",
			input:    "Hello Привет",
			expected: 4,
		},
		{
			name:     "Test#7: String with punctuation and spaces",
			input:    "A quick brown fox jumps over the lazy dog.",
			expected: 12,
		},
		{
			name:     "String with uppercase vowels",
			input:    "AEIOU",
			expected: 5,
		},
		{
			name:     "String with mixed case vowels",
			input:    "AeIoU",
			expected: 5,
		},
		{
			name:     "String with no vowels in Cyrillic",
			input:    "Бржднт, Мнп!",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountVowels(tt.input)
			if got != tt.expected {
				t.Errorf("CountVowels(%q) = %d; want %d", tt.input, got, tt.expected)
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
	expected := "3\n3\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
