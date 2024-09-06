package main

import (
	"bytes"
	"os"
	"testing"
)

func TestFilterWords(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		censorMap map[string]string
		expected  string
	}{
		{
			name:  "Basic censorship",
			input: "Купи биткоин сейчас! Биткоин очень выгоден!",
			censorMap: map[string]string{
				"биткоин": "фрукты",
			},
			expected: "Купи фрукты сейчас! Фрукты очень выгоден!",
		},
		{
			name:  "Case insensitive censorship",
			input: "Купи Биткоин сейчас! Биткоин очень выгоден!",
			censorMap: map[string]string{
				"биткоин": "фрукты",
			},
			expected: "Купи Фрукты сейчас! Фрукты очень выгоден!",
		},
		{
			name:  "Multiple words censorship",
			input: "Купи биткоин и эфир сейчас! Биткоин и эфир - наше будущее!",
			censorMap: map[string]string{
				"биткоин": "фрукты",
				"эфир":    "овощи",
			},
			expected: "Купи фрукты и овощи сейчас! Фрукты и овощи - наше будущее!",
		},
		{
			name:      "No censorship",
			input:     "Купи биткоин сейчас!",
			censorMap: map[string]string{},
			expected:  "Купи биткоин сейчас!",
		},
		{
			name:  "Word uniqueness",
			input: "Внимание! Внимание! Покупай срочно срочно крипту только у нас!",
			censorMap: map[string]string{
				"крипту": "фрукты",
			},
			expected: "Внимание! Покупай срочно фрукты только у нас!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := filterWords(tt.input, tt.censorMap)
			if actual != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, actual)
			}
		})
	}
}

func TestCheckUpper(t *testing.T) {
	tests := []struct {
		old, new string
		expected string
	}{
		{"Hello", "world", "World"}, // Test for capitalizing first letter
		{"hello", "world", "world"}, // Test for not capitalizing first letter
		{"", "world", "world"},      // Test with empty old string
		{"Hello", "", ""},           // Test with empty new string
		{"", "", ""},                // Test with both strings empty
	}

	for _, tt := range tests {
		t.Run(tt.old+"_"+tt.new, func(t *testing.T) {
			if result := CheckUpper(tt.old, tt.new); result != tt.expected {
				t.Errorf("CheckUpper(%q, %q) = %q; want %q", tt.old, tt.new, result, tt.expected)
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
	expected := "Внимание! Покупай срочно фрукты только у нас! Фрукты лайткоин фрукты по низким ценам! Беги, успевай стать финансово независимым с помощью фруктов! Фрукты будущее финансового мира!\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
