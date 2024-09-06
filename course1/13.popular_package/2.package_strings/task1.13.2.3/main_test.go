package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name   string
		length int
		want   int
	}{
		{
			name:   "Test#1",
			length: 5,
			want:   5,
		},
		{
			name:   "Test#2",
			length: 0,
			want:   0,
		},
		{
			name:   "Test#3",
			length: 10,
			want:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomString(tt.length); len([]rune(got)) != tt.want {
				t.Errorf("GenerateRandomString() = %v, want %v", got, tt.want)
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
	expected := 10

	if len([]rune(result)) != expected {
		t.Errorf("got %d, want %d", len([]rune(result)), expected)
	}
}
