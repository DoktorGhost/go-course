package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMaxDifference(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		{
			name:    "Test #1: nil",
			numbers: []int{},
			want:    0,
		},
		{
			name:    "Test #2: one element",
			numbers: []int{1},
			want:    0,
		},
		{
			name:    "Test #3: max=min",
			numbers: []int{1, 1},
			want:    0,
		},
		{
			name:    "Test #4: normal",
			numbers: []int{36, 31, 3, 52, 5, 1, 101},
			want:    100,
		},
		{
			name:    "Test #5: negative numbers",
			numbers: []int{-10, -5, -3, -7},
			want:    7,
		},
		{
			name:    "Test #6: all same values",
			numbers: []int{7, 7, 7, 7},
			want:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxDifference(tt.numbers); got != tt.want {
				t.Errorf("MaxDifference() = %v, want %v", got, tt.want)
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
	expected := "9\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
