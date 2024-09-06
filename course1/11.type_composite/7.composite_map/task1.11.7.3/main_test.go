package main

import (
	"bytes"
	"os"
	"testing"
)

func Test_createUniqueText(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Test#1: normal",
			arg:  "bar bar count ssh ssh yes no",
			want: "bar count ssh yes no",
		},
		{
			name: "Test#2: nil",
			arg:  "",
			want: "",
		},
		{
			name: "Test#3: normal",
			arg:  "bar",
			want: "bar",
		},
		{
			name: "Test#4: normal",
			arg:  "bar bar bar bar bar bar bar bar",
			want: "bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createUniqueText(tt.arg); got != tt.want {
				t.Errorf("createUniqueText() = %v, want %v", got, tt.want)
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
	expected := "bar foo baz\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
