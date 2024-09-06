package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func Test_countRussianLetters(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "test#1",
			s:    "Привет Медвед!",
			want: 8,
		},
		{
			name: "test#2",
			s:    " ",
			want: 0,
		},
		{
			name: "test#3",
			s:    "22 32 432 4252",
			want: 0,
		},
		{
			name: "test#4",
			s:    "Hello my friends, Александр!",
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRussianLetters(tt.s); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("countRussianLetters() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func Test_isRussianLetter(t *testing.T) {

	tests := []struct {
		name string
		char rune
		want bool
	}{
		{
			name: "Test#1",
			char: 'А',
			want: true,
		},
		{
			name: "Test#2",
			char: 'A',
			want: false,
		},
		{
			name: "Test#3",
			char: 'а',
			want: true,
		},
		{
			name: "Test#4",
			char: 'a',
			want: false,
		},
		{
			name: "Test#5",
			char: ' ',
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRussianLetter(tt.char); got != tt.want {
				t.Errorf("isRussianLetter() = %v, want %v", got, tt.want)
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
	lens := len(stdout.String())
	expected := 42

	if lens != expected {
		t.Errorf("got %d, want %d", lens, expected)
	}
}
