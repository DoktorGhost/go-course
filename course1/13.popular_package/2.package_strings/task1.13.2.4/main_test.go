package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func Test_generateActivationKey(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Test#1",
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateActivationKey(); len([]rune(got)) != tt.want {
				t.Errorf("generateActivationKey() = %v, want %v", got, tt.want)
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
	result := strings.Trim(stdout.String(), "\n")
	expected := 19

	if len([]rune(result)) != expected {
		t.Errorf("got %d, want %d", len([]rune(result)), expected)
	}
}
