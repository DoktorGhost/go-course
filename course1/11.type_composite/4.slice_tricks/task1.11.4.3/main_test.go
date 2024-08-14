package main

import (
	"reflect"
	"testing"
)

func TestRemoveExtraMemory(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Test#1: normal",
			args: make([]int, 5, 10),
			want: 5,
		},
		{
			name: "Test#2: zero",
			args: make([]int, 0, 10),
			want: 0,
		},
		{
			name: "Test#2: zero",
			args: make([]int, 10, 10),
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveExtraMemory(tt.args); !reflect.DeepEqual(cap(got), tt.want) {
				t.Errorf("RemoveExtraMemory() = %v, want %v", cap(got), tt.want)
			}
		})
	}
}
