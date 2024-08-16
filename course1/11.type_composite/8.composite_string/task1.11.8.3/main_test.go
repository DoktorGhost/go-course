package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func Test_getBytes(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []byte
	}{
		{
			name: "Test#1: en",
			args: "Hello",
			want: []byte("Hello"),
		},
		{
			name: "Test#2: ру",
			args: "Привет",
			want: []byte("Привет"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBytes(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRunes(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []rune
	}{
		{
			name: "Test#1: en",
			args: "Hello",
			want: []rune("Hello"),
		},
		{
			name: "Test#2: ру",
			args: "Привет",
			want: []rune("Привет"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRunes(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRunes() = %v, want %v", got, tt.want)
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
	expected := "[72 101 108 108 111 44 32 1052 1048 1056 33]\n[72 101 108 108 111 44 32 208 156 208 152 208 160 33]\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
