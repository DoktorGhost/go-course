package main

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{
			name: "Test 1: factorial(0)",
			arg:  0,
			want: 1,
		},
		{
			name: "Test 2: factorial(1)",
			arg:  1,
			want: 1,
		},
		{
			name: "Test 3: factorial(5)",
			arg:  5,
			want: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.arg); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
