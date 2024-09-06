package main

import (
	"bytes"
	"os"
	"testing"
)

func Test_getVariableType(t *testing.T) {
	type args struct {
		variable interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test#1",
			args: args{
				variable: 12,
			},
			want: "int",
		},
		{
			name: "Test#2",
			args: args{
				variable: "12",
			},
			want: "string",
		},
		{
			name: "Test#3",
			args: args{
				variable: true,
			},
			want: "bool",
		},
		{
			name: "Test#4",
			args: args{
				variable: 2.56,
			},
			want: "float64",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getVariableType(tt.args.variable); got != tt.want {
				t.Errorf("getVariableType() = %v, want %v", got, tt.want)
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
	expected := "int\nstring\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
