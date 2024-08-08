package main

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{
			name: "Test 1: Fibonacci(5)",
			arg:  5,
			want: 5,
		},
		{
			name: "Test 2: Fibonacci(6)",
			arg:  6,
			want: 8,
		},
		{
			name: "Test 3: Fibonacci(7)",
			arg:  7,
			want: 13,
		},
		{
			name: "Test 4: Fibonacci(8)",
			arg:  8,
			want: 21,
		},
		{
			name: "Test 5: Fibonacci(9)",
			arg:  9,
			want: 34,
		},
		{
			name: "Test 6: Fibonacci(10)",
			arg:  10,
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.arg); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFibonacci(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Fibonacci(10)
	}
}
