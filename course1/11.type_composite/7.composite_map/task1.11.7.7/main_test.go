package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func Test_getUniqueUsers(t *testing.T) {
	tests := []struct {
		name string
		args []User
		want []User
	}{
		{
			name: "Test#1: no duplicates",
			args: []User{
				{Nickname: "John", Age: 25, Email: "john@gmail.com"},
				{Nickname: "Jane", Age: 25, Email: "jane@gmail.com"},
				{Nickname: "John2", Age: 48, Email: "john24@gmail.com"},
			},
			want: []User{
				{Nickname: "John", Age: 25, Email: "john@gmail.com"},
				{Nickname: "Jane", Age: 25, Email: "jane@gmail.com"},
				{Nickname: "John2", Age: 48, Email: "john24@gmail.com"},
			},
		},
		{
			name: "Test#2: normal",
			args: []User{
				{Nickname: "John", Age: 25, Email: "john@gmail.com"},
				{Nickname: "Jane", Age: 25, Email: "jane@gmail.com"},
				{Nickname: "John", Age: 30, Email: "john2@gmail.com"},
			},
			want: []User{
				{Nickname: "John", Age: 25, Email: "john@gmail.com"},
				{Nickname: "Jane", Age: 25, Email: "jane@gmail.com"},
			},
		},
		{
			name: "Test#3: Empty list",
			args: []User{},
			want: []User{},
		},
		{
			name: "Test#4: All duplicates",
			args: []User{
				{Nickname: "John", Age: 25, Email: "john@gmail.com"},
				{Nickname: "John", Age: 30, Email: "john2@gmail.com"},
				{Nickname: "John", Age: 35, Email: "john3@gmail.com"},
			},
			want: []User{
				{Nickname: "John", Age: 25, Email: "john@gmail.com"},
			},
		},
		{
			name: "Test#5: Single user",
			args: []User{
				{Nickname: "John", Age: 25, Email: "john@gmail.com"},
			},
			want: []User{
				{Nickname: "John", Age: 25, Email: "john@gmail.com"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUniqueUsers(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUniqueUsers() = %v, want %v", got, tt.want)
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
	expected := "[{John 25 john@gmail.com} {Jane 25 jane@gmail.com} {John2 48 john24@gmail.com}]\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
