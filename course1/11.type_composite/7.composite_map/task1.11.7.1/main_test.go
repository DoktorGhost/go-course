package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestCountWordOccurences(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:  "Basic case",
			input: "Lorem ipsum, dolor sit amet! Consectetur adipiscing elit. Ipsum, lorem!",
			expected: map[string]int{
				"lorem":       2,
				"ipsum":       2,
				"dolor":       1,
				"sit":         1,
				"amet":        1,
				"consectetur": 1,
				"adipiscing":  1,
				"elit":        1,
			},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: map[string]int{
				// Expect an empty map
			},
		},
		{
			name:     "String with only punctuation",
			input:    "!!!,,,;;;...",
			expected: map[string]int{
				// Expect an empty map
			},
		},
		{
			name:  "Case insensitive",
			input: "Go gO GO go!",
			expected: map[string]int{
				"go": 4,
			},
		},
		{
			name:  "Numbers included",
			input: "123 123 go 456",
			expected: map[string]int{
				"123": 2,
				"go":  1,
				"456": 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countWordOccurences(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("For input '%s', expected %v but got %v", tt.input, tt.expected, actual)
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

	expected := 79

	if (len(stdout.String())) != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
