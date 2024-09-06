package main

import (
	"bytes"
	"os"
	"testing"
)

func Test_generateMathString(t *testing.T) {
	type args struct {
		operands []int
		operator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test#1",
			args: args{
				operands: []int{1, 2, 3, 4, 5},
				operator: "+",
			},
			want: "1+2+3+4+5=15",
		},
		{
			name: "test#2",
			args: args{
				operands: []int{1},
				operator: "+",
			},
			want: "передан один операнд",
		},
		{
			name: "test#3",
			args: args{
				operands: []int{1, 0, 2},
				operator: "/",
			},
			want: "делить на 0 нельзя",
		},
		{
			name: "test#4",
			args: args{
				operands: []int{100, 2},
				operator: "/",
			},
			want: "100/2=50",
		},
		{
			name: "test#5",
			args: args{
				operands: []int{100, 2, 50},
				operator: "-",
			},
			want: "100-2-50=48",
		},
		{
			name: "test#6",
			args: args{
				operands: []int{-100, 2, 50},
				operator: "-",
			},
			want: "-100-2-50=-152",
		},
		{
			name: "test#7",
			args: args{
				operands: []int{1, 2, 3},
				operator: "*",
			},
			want: "1*2*3=6",
		},
		{
			name: "test#8",
			args: args{
				operands: []int{1, 2, 3},
				operator: "",
			},
			want: "не передан аргумент",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMathString(tt.args.operands, tt.args.operator); got != tt.want {
				t.Errorf("generateMathString() = %v, want %v", got, tt.want)
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
	expected := "2+4+6=12\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
