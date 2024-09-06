package main

import (
	"bytes"
	"os"
	"testing"
)

func Test_isValidEmail(t *testing.T) {

	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{
			name:  "Test#1",
			email: "test@test.com",
			want:  true,
		},
		{
			name:  "Test#2",
			email: "testtest.com",
			want:  false,
		},
		{
			name:  "Test#3",
			email: "",
			want:  false,
		},
		{
			name:  "Test#4",
			email: "1@2.3",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidEmail(tt.email); got != tt.want {
				t.Errorf("isValidEmail() = %v, want %v", got, tt.want)
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
	expected := "test@example.com является валидным e-mail адресом\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
