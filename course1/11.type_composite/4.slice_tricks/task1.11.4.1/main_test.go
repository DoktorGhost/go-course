package main

import (
	"reflect"
	"testing"
)

func TestCut(t *testing.T) {
	type args struct {
		xs    []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test#1: normal",
			args: args{
				xs:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				start: 1,
				end:   5,
			},
			want: []int{2, 3, 4, 5, 6},
		},
		{
			name: "Test#2: start > finish",
			args: args{
				xs:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				start: 5,
				end:   2,
			},
			want: []int{},
		},
		{
			name: "Test#3: start = finish",
			args: args{
				xs:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				start: 2,
				end:   2,
			},
			want: []int{3},
		},
		{
			name: "Test#4: start < 0",
			args: args{
				xs:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				start: -2,
				end:   2,
			},
			want: []int{},
		},
		{
			name: "Test#5: finish < 0",
			args: args{
				xs:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				start: 2,
				end:   -2,
			},
			want: []int{},
		},
		{
			name: "Test#5: finish < 0 && start < 0",
			args: args{
				xs:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				start: -2,
				end:   -2,
			},
			want: []int{},
		},
		{
			name: "Test#5: finish < 0 && start < 0",
			args: args{
				xs:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				start: 1,
				end:   9,
			},
			want: []int{2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cut(tt.args.xs, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cut() = %v, want %v", got, tt.want)
			}
		})
	}
}
