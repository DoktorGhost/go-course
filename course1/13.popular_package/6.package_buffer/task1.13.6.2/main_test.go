package main

import (
	"bytes"
	"os"
	"testing"
)

// Тестовая функция для getScanner
func TestGetScanner(t *testing.T) {
	tests := []struct {
		input  string
		output []string
	}{
		{
			input:  "Hello\nWorld",
			output: []string{"Hello", "World"},
		},
		{
			input:  "SingleLine",
			output: []string{"SingleLine"},
		},
		{
			input:  "LeadingNewline\nTrailingNewline\n",
			output: []string{"LeadingNewline", "TrailingNewline"},
		},
		{
			input:  "Line1\nLine2\nLine3",
			output: []string{"Line1", "Line2", "Line3"},
		},
	}

	for _, test := range tests {
		buffer := bytes.NewBufferString(test.input)
		scanner := getScanner(buffer)

		var result []string
		for scanner.Scan() {
			result = append(result, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			t.Fatalf("Scan error: %v", err)
		}

		if len(result) != len(test.output) {
			t.Errorf("Expected %d lines, got %d", len(test.output), len(result))
			continue
		}

		for i, line := range test.output {
			if result[i] != line {
				t.Errorf("Line %d: expected %q, got %q", i, line, result[i])
			}
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
	expected := "Hello\n,\n World!\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
