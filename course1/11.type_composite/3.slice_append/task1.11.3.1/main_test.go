package main

import (
	"reflect"
	"testing"
)

func Test_appendInt(t *testing.T) {
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
			name: "test#1: nil slice",
			args: args{
				xs: []int{},
				x:  []int{},
			},
			want: []int{},
		},
		{
			name: "test#2: nil slice added",
			args: args{
				xs: []int{1, 2, 3},
				x:  []int{},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "test#3: normal",
			args: args{
				xs: []int{1, 2, 3},
				x:  []int{1, 2, 3},
			},
			want: []int{1, 2, 3, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendInt(tt.args.xs, tt.args.x...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
