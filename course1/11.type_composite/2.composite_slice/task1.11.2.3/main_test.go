package main

import (
	"bytes"
	"os"
	"testing"
)

func Test_bitwiseXOR(t *testing.T) {
	type args struct {
		n   int
		res int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test #1: 0 ^ 1",
			args: args{
				n:   0,
				res: 1,
			},
			want: 1,
		},
		{
			name: "Test #2: 1 ^ 1",
			args: args{
				n:   1,
				res: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitwiseXOR(tt.args.n, tt.args.res); got != tt.want {
				t.Errorf("bitwiseXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSingleNumber(t *testing.T) {

	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		{
			name:    "Test #1: 0",
			numbers: []int{0},
			want:    0,
		},
		{
			name:    "Test #2: not",
			numbers: []int{1, 1},
			want:    0,
		},
		{
			name:    "Test #3: normal",
			numbers: []int{1, 1, 3},
			want:    3,
		},
		{
			name:    "Test #3: normal 2",
			numbers: []int{0, 0, 5},
			want:    5,
		},
		{
			name:    "Test #4: negative number",
			numbers: []int{-1, -2, -1},
			want:    -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSingleNumber(tt.numbers); got != tt.want {
				t.Errorf("findSingleNumber() = %v, want %v", got, tt.want)
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
	expected := "5\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
