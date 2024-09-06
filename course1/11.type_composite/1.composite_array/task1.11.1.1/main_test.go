package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func Test_average(t *testing.T) {
	tests := []struct {
		name string
		args [8]int
		want float64
	}{
		{
			name: "Test #1: nil",
			args: [8]int{},
			want: 0.0,
		},
		{
			name: "Test #2: normal",
			args: [8]int{1, 2, 3, 4, 5, 6, 7, 8},
			want: 4.5,
		},
		{
			name: "Test #3: one element",
			args: [8]int{1},
			want: 0.125,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := average(tt.args); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_averageFloat(t *testing.T) {
	tests := []struct {
		name string
		args [8]float64
		want float64
	}{
		{
			name: "Test #1: nil",
			args: [8]float64{},
			want: 0.0,
		},
		{
			name: "Test #2: normal",
			args: [8]float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5},
			want: 5.0,
		},
		{
			name: "Test #3: one element",
			args: [8]float64{1.5},
			want: 0.1875,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageFloat(tt.args); got != tt.want {
				t.Errorf("averageFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		args [8]int
		want [8]int
	}{
		{
			name: "Test #1: nil",
			args: [8]int{},
			want: [8]int{},
		},
		{
			name: "Test #2: normal",
			args: [8]int{1, 2, 3, 4, 5, 6, 7, 8},
			want: [8]int{8, 7, 6, 5, 4, 3, 2, 1},
		},

		{
			name: "Test #3: one element",
			args: [8]int{1},
			want: [8]int{0, 0, 0, 0, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sum(t *testing.T) {
	tests := []struct {
		name string
		args [8]int
		want int
	}{
		{
			name: "Test #1: nil",
			args: [8]int{},
			want: 0,
		},
		{
			name: "Test #2: normal",
			args: [8]int{1, 2, 3, 4, 5, 6, 7, 8},
			want: 36,
		},

		{
			name: "Test #3: one element",
			args: [8]int{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
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
	expected := "36\n4.5\n5\n[8 7 6 5 4 3 2 1]\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
