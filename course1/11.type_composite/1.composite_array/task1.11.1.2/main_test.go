package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func Test_sortAscFloat(t *testing.T) {

	tests := []struct {
		name string
		args [8]float64
		want [8]float64
	}{
		{
			name: "Test #1: nil",
			args: [8]float64{},
			want: [8]float64{},
		},
		{
			name: "Test #2: normal",
			args: [8]float64{1.1, 2.2, 3.3, 4.4, 1.1, 2.2, 3.3, 4.4},
			want: [8]float64{1.1, 1.1, 2.2, 2.2, 3.3, 3.3, 4.4, 4.4},
		},
		{
			name: "Test #3: zero",
			args: [8]float64{1.1, 2.2, 3.3, 4.4},
			want: [8]float64{0, 0.0, 0, 0, 1.1, 2.2, 3.3, 4.4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortAscFloat(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortAscFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortAscInt(t *testing.T) {

	tests := []struct {
		name string
		args [8]int
		want [8]int
	}{
		{
			name: "Test #1: nil",
			args: [8]int{},
			want: [8]int{},
		},
		{
			name: "Test #2: normal",
			args: [8]int{2, 4, 7, 6, 3, 1, 2, 8},
			want: [8]int{1, 2, 2, 3, 4, 6, 7, 8},
		},
		{
			name: "Test #3: zero",
			args: [8]int{8, 1},
			want: [8]int{0, 0, 0, 0, 0, 0, 1, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortAscInt(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortAscInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortDescFloat(t *testing.T) {
	type args struct {
		floatArr [8]float64
	}
	tests := []struct {
		name string
		args [8]float64
		want [8]float64
	}{
		{
			name: "Test #1: nil",
			args: [8]float64{},
			want: [8]float64{},
		},
		{
			name: "Test #2: normal",
			args: [8]float64{1.1, 2.2, 3.3, 4.4, 1.1, 2.2, 3.3, 4.4},
			want: [8]float64{4.4, 4.4, 3.3, 3.3, 2.2, 2.2, 1.1, 1.1},
		},
		{
			name: "Test #3: zero",
			args: [8]float64{1.1, 2.2, 3.3, 4.4},
			want: [8]float64{4.4, 3.3, 2.2, 1.1, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortDescFloat(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortDescFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortDescInt(t *testing.T) {
	tests := []struct {
		name string
		args [8]int
		want [8]int
	}{
		{
			name: "Test #1: nil",
			args: [8]int{},
			want: [8]int{},
		},
		{
			name: "Test #2: normal",
			args: [8]int{2, 4, 7, 6, 3, 1, 2, 8},
			want: [8]int{8, 7, 6, 4, 3, 2, 2, 1},
		},
		{
			name: "Test #3: zero",
			args: [8]int{8, 1},
			want: [8]int{8, 1, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortDescInt(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortDescInt() = %v, want %v", got, tt.want)
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
	expected := "Sorted Int Array (Descending): [9 8 7 5 4 3 2 1]\nSorted Int Array (Ascending): [1 2 3 4 5 7 8 9]\n" +
		"Sorted Float Array (Descending): [9.9 8.8 7.7 5.5 4.4 3.3 2.2 1.1]\nSorted Float Array (Ascending): [1.1 2.2 3.3 4.4 5.5 7.7 8.8 9.9]\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
