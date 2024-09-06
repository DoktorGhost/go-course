package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"unsafe"
)

func TestGetStringHeaders(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedLen int
	}{
		{
			name:        "Normal string",
			input:       "Hello, World!",
			expectedLen: len("Hello, World!"),
		},
		{
			name:        "Empty string",
			input:       "",
			expectedLen: 0,
		},
		{
			name:        "Single character string",
			input:       "A",
			expectedLen: 1,
		},
		{
			name:        "Long string",
			input:       "This is a much longer string to test the edge cases.",
			expectedLen: len("This is a much longer string to test the edge cases."),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header := getStringHeaders(tt.input)

			// Validate the Len field
			if header.Len != tt.expectedLen {
				t.Errorf("Expected Len: %v, but got: %v", tt.expectedLen, header.Len)
			}

			// Optional: Validate the Data field if you can derive a valid reference
			// Instead of comparing Data directly, you could validate via reflection or manual check
			if tt.expectedLen > 0 {
				// Ensure that Data points to a valid memory location, as exact matches are not reliable
				dataPtr := unsafe.Pointer(uintptr(header.Data))
				if dataPtr == nil {
					t.Errorf("Expected Data to be non-nil, but got nil.")
				}
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
	expectedPrefix := "Data: "
	expectedLen := "Len: 13\n"

	output := stdout.String()

	// Проверяем, что вывод содержит ожидаемую длину
	if !strings.HasSuffix(output, expectedLen) {
		t.Errorf("expected output to end with %q, but got %q", expectedLen, output)
	}

	// Проверяем, что вывод содержит префикс Data: и длину
	if !strings.HasPrefix(output, expectedPrefix) {
		t.Errorf("expected output to start with %q, but got %q", expectedPrefix, output)
	}
}
