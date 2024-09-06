package main

import (
	"bytes"
	"os"
	"testing"
)

func Test_countUniqueUTF8Chars(t *testing.T) {

	tests := []struct {
		name string
		arg  string
		want int
	}{
		{
			name: "Test#1",
			arg:  "Hello",
			want: 4,
		},
		{
			name: "Test#2",
			arg:  "Привет",
			want: 6,
		},
		{
			name: "Test#3",
			arg:  "",
			want: 0,
		},
		{
			name: "Test#4",
			arg:  " ",
			want: 1,
		},
		{
			name: "Test#4",
			arg:  ".,/!",
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUniqueUTF8Chars(tt.arg); got != tt.want {
				t.Errorf("countUniqueUTF8Chars() = %v, want %v", got, tt.want)
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
	expected := "7\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
