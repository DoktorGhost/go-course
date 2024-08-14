package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func Test_getSubSlice(t *testing.T) {
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
			name: "test #1: nil",
			args: args{
				xs:    []int{},
				start: 0,
				end:   0,
			},
			want: []int{},
		},
		{
			name: "test #2: normal",
			args: args{
				xs:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				start: 1,
				end:   4,
			},
			want: []int{2, 3, 4},
		},
		{
			name: "test #3: bad start",
			args: args{
				xs:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				start: 11,
				end:   4,
			},
			want: []int{},
		},
		{
			name: "test #3: bad finish",
			args: args{
				xs:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				start: 1,
				end:   12,
			},
			want: []int{},
		},
		{
			name: "test #3: start > finish",
			args: args{
				xs:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				start: 5,
				end:   3,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSubSlice(tt.args.xs, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSubSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	var stdout bytes.Buffer
	stdout.ReadFrom(r)
	expected := "[3 4 5 6]\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
