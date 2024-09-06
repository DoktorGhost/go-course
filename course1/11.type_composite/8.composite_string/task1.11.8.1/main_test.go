package main

import "testing"

func Test_countBytes(t *testing.T) {

	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "Test#1 countBytes",
			args: "Привет, мир!",
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBytes(tt.args); got != tt.want {
				t.Errorf("countBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countSymbol(t *testing.T) {

	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "Test#1 countSymbol",
			args: "Привет, мир!",
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSymbol(tt.args); got != tt.want {
				t.Errorf("countSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
