package main

import (
	"reflect"
	"testing"
)

func TestInsertToStart(t *testing.T) {
	type args struct {
		xs []int
		x  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test#1: normal",
			args: args{
				xs: []int{1, 2, 3},
				x:  []int{4, 5, 6},
			},
			want: []int{4, 5, 6, 1, 2, 3},
		},
		{
			name: "Test#2: nil",
			args: args{
				xs: []int{},
				x:  []int{4, 5, 6},
			},
			want: []int{4, 5, 6},
		},
		{
			name: "Test#3: nil2",
			args: args{
				xs: []int{4, 5, 6},
				x:  []int{},
			},
			want: []int{4, 5, 6},
		},
		{
			name: "Test#3: nil3",
			args: args{
				xs: []int{},
				x:  []int{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertToStart(tt.args.xs, tt.args.x...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertToStart() = %v, want %v", got, tt.want)
			}
		})
	}
}
