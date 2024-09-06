package main

import (
	"reflect"
	"testing"
)

func TestFilterDividers(t *testing.T) {
	type args struct {
		xs      []int
		divider int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test#1: normal",
			args: args{
				xs:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				divider: 2,
			},
			want: []int{2, 4, 6, 8, 10},
		},
		{
			name: "Test#2: negativ",
			args: args{
				xs:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				divider: -2,
			},
			want: []int{2, 4, 6, 8, 10},
		},
		{
			name: "Test#3: nil",
			args: args{
				xs:      []int{},
				divider: -2,
			},
			want: []int{},
		},
		{
			name: "Test#4: zero",
			args: args{
				xs:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				divider: 0,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterDividers(tt.args.xs, tt.args.divider); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterDividers() = %v, want %v", got, tt.want)
			}
		})
	}
}
