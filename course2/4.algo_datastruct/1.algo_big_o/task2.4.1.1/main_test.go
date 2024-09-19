package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"testing"
)

func Test_compareWhichFactorialIsFaster(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "false",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareWhichFactorialIsFaster(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compareWhichFactorialIsFaster() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_factorialIterative(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{
			name: "0",
			arg:  0,
			want: 0,
		},
		{
			name: "1",
			arg:  1,
			want: 1,
		},
		{
			name: "3",
			arg:  3,
			want: 6,
		},
		{
			name: "5",
			arg:  5,
			want: 120,
		},
		{
			name: "7",
			arg:  7,
			want: 5040,
		},
		{
			name: "10",
			arg:  10,
			want: 3628800,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorialIterative(tt.arg); got != tt.want {
				t.Errorf("factorialIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_factorialRecursive(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{
			name: "0",
			arg:  0,
			want: 0,
		},
		{
			name: "1",
			arg:  1,
			want: 1,
		},
		{
			name: "3",
			arg:  3,
			want: 6,
		},
		{
			name: "5",
			arg:  5,
			want: 120,
		},
		{
			name: "7",
			arg:  7,
			want: 5040,
		},
		{
			name: "10",
			arg:  10,
			want: 3628800,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorialRecursive(tt.arg); got != tt.want {
				t.Errorf("factorialRecursive() = %v, want %v", got, tt.want)
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

	expected := fmt.Sprintf("Go version: %v\nGo OS/Arch: %v / %v\nWhich factorial is faster?\nfalse\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
