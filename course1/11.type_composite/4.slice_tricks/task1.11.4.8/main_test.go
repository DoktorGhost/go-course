package main

import (
	"reflect"
	"testing"
)

func TestShift(t *testing.T) {
	tests := []struct {
		name  string
		args  []int
		want  int
		want1 []int
	}{
		{
			name:  "Test#1: normal",
			args:  []int{1, 2, 3, 4, 5},
			want:  1,
			want1: []int{5, 1, 2, 3, 4},
		},
		{
			name:  "Test#2: nil",
			args:  []int{},
			want:  0,
			want1: []int{},
		},
		{
			name:  "Test#3: one element",
			args:  []int{10},
			want:  10,
			want1: []int{10},
		},
		{
			name:  "Test#4: two elements",
			args:  []int{10, 11},
			want:  10,
			want1: []int{11, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Shift(tt.args)
			if got != tt.want {
				t.Errorf("Shift() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Shift() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
