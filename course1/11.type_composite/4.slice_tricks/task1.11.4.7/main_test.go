package main

import (
	"reflect"
	"testing"
)

func TestPop(t *testing.T) {

	tests := []struct {
		name  string
		args  []int
		want  int
		want1 []int
	}{
		{
			name:  "Test#1: normal",
			args:  []int{1, 2, 3},
			want:  1,
			want1: []int{2, 3},
		},
		{
			name:  "Test#2: nil",
			args:  []int{},
			want:  0,
			want1: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Pop(tt.args)
			if got != tt.want {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
