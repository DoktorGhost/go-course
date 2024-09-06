package main

import (
	"reflect"
	"testing"
)

func TestInsertAfterIDX(t *testing.T) {
	type args struct {
		xs  []int
		idx int
		x   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test#1: normal",
			args: args{
				xs:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				idx: 1,
				x:   []int{1, 2, 3},
			},
			want: []int{1, 2, 1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "test#2: idx < 0",
			args: args{
				xs:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				idx: -1,
				x:   []int{1, 2, 3},
			},
			want: []int{},
		},
		{
			name: "test#3: idx > len",
			args: args{
				xs:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				idx: 15,
				x:   []int{1, 2, 3},
			},
			want: []int{},
		},
		{
			name: "test#4: idx = len-1",
			args: args{
				xs:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				idx: 9,
				x:   []int{1, 2, 3},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3},
		},
		{
			name: "test#5: idx = 0",
			args: args{
				xs:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				idx: 0,
				x:   []int{1, 2, 3},
			},
			want: []int{1, 1, 2, 3, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertAfterIDX(tt.args.xs, tt.args.idx, tt.args.x...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertAfterIDX() = %v, want %v", got, tt.want)
			}
		})
	}
}
