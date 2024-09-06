package main

import (
	"reflect"
	"testing"
)

func TestRemoveIDX(t *testing.T) {
	type args struct {
		xs  []int
		idx int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test#1: normal",
			args: args{
				xs:  []int{1, 2, 3, 4},
				idx: 0,
			},
			want: []int{2, 3, 4},
		},
		{
			name: "Test#2: nil",
			args: args{
				xs:  []int{},
				idx: 0,
			},
			want: []int{},
		},
		{
			name: "Test#3: idx > len",
			args: args{
				xs:  []int{1, 2, 3, 4},
				idx: 8,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Test#4: negativ idx",
			args: args{
				xs:  []int{1, 2, 3, 4},
				idx: -8,
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveIDX(tt.args.xs, tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIDX() = %v, want %v", got, tt.want)
			}
		})
	}
}
