package main

import (
	"reflect"
	"testing"
)

func Test_mergeMaps(t *testing.T) {
	type args struct {
		map1 map[string]int
		map2 map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Test#1: normal",
			args: args{
				map1: map[string]int{"a": 1, "b": 2},
				map2: map[string]int{"c": 1, "d": 2},
			},
			want: map[string]int{"a": 1, "b": 2, "c": 1, "d": 2},
		},
		{
			name: "Test#2: key=key",
			args: args{
				map1: map[string]int{"a": 1, "b": 2},
				map2: map[string]int{"a": 1, "d": 2},
			},
			want: map[string]int{"a": 1, "b": 2, "d": 2},
		},
		{
			name: "Test#3:nil map",
			args: args{
				map1: map[string]int{"a": 1, "b": 2},
				map2: map[string]int{},
			},
			want: map[string]int{"a": 1, "b": 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeMaps(tt.args.map1, tt.args.map2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
