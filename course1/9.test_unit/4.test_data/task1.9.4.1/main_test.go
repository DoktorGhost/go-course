package main

import (
	"reflect"
	"testing"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		name string
		args []float64
		want float64
	}{
		{
			name: "Test 1: normal case",
			args: []float64{1, 2, 3},
			want: 2,
		},
		{
			name: "Test 2: nil slice",
			args: []float64{},
			want: 0,
		},
		{
			name: "Test 3: negative case",
			args: []float64{-1, -2, -3},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := average(tt.args); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateTestData(t *testing.T) {
	size := 5
	slice1 := generateTestData(size)
	slice2 := generateTestData(size)

	if !reflect.DeepEqual(len(slice1), len(slice2)) {
		t.Errorf("%v != %v", len(slice1), len(slice2))
	}
}
