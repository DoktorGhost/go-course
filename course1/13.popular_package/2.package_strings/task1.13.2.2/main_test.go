package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestFilterComments(t *testing.T) {

	tests := []struct {
		name  string
		users []User
		want  []User
	}{
		{
			name: "Test#1",
			users: []User{
				{
					Name: "Betty",
					Comments: []Comment{
						{Message: "good Comment 1"},
						{Message: "BaD CoMmEnT 2"},
						{Message: "Bad Comment 3"},
						{Message: "Use camelCase please"},
					},
				},
				{
					Name: "Jhon",
					Comments: []Comment{
						{Message: "Good Comment 1"},
						{Message: "Good Comment 2"},
						{Message: "Good Comment 3"},
						{Message: "Bad Comment 4"},
					},
				},
			},
			want: []User{
				{
					Name: "Betty",
					Comments: []Comment{
						{Message: "good Comment 1"},
						{Message: "Use camelCase please"},
					},
				},
				{
					Name: "Jhon",
					Comments: []Comment{
						{Message: "Good Comment 1"},
						{Message: "Good Comment 2"},
						{Message: "Good Comment 3"},
					},
				},
			},
		},
		{
			name: "Test#2",
			users: []User{
				{
					Name:     "Betty",
					Comments: []Comment{},
				},
				{
					Name:     "Jhon",
					Comments: []Comment{},
				},
			},
			want: []User{
				{
					Name:     "Betty",
					Comments: []Comment{},
				},
				{
					Name:     "Jhon",
					Comments: []Comment{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterComments(tt.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBadComments(t *testing.T) {
	tests := []struct {
		name string
		user User
		want []Comment
	}{
		{
			name: "Test#1",
			user: User{
				Name: "Betty",
				Comments: []Comment{
					{Message: "good Comment 1"},
					{Message: "baD Comment 2"},
					{Message: "good Comment 3"},
					{Message: "BAD Comment 4"},
				},
			},
			want: []Comment{
				{Message: "baD Comment 2"},
				{Message: "BAD Comment 4"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBadComments(tt.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBadComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGoodComments(t *testing.T) {

	tests := []struct {
		name string
		user User
		want []Comment
	}{
		{
			name: "Test#1",
			user: User{
				Name: "Betty",
				Comments: []Comment{
					{Message: "good Comment 1"},
					{Message: "baD Comment 2"},
					{Message: "good Comment 3"},
					{Message: "BAD Comment 4"},
				},
			},
			want: []Comment{
				{Message: "good Comment 1"},
				{Message: "good Comment 3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGoodComments(tt.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGoodComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBadComment(t *testing.T) {

	tests := []struct {
		name    string
		comment string
		want    bool
	}{
		{
			name:    "Test#1",
			comment: "good Comment 1",
			want:    false,
		},
		{
			name:    "Test#2",
			comment: "bad Comment 1",
			want:    true,
		},
		{
			name:    "Test#3",
			comment: "",
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBadComment(tt.comment); got != tt.want {
				t.Errorf("IsBadComment() = %v, want %v", got, tt.want)
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
	expected := "[{Betty [{good Comment 1} {Use camelCase please}]} {Jhon [{Good Comment 1} {Good Comment 2} {Good Comment 3}]}]\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
