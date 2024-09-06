package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func Test_getUsersFromJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    []User
		wantErr bool
	}{
		{
			name: "Test#1",
			data: []byte(`
				[ 
					{
						"name": "Jhon",
						"age": 30,
						"comments": [
							{"text": "Great post!"},
							{"text": "I agree"}
						]
					}
				]`),
			want: []User{
				{
					Name: "Jhon",
					Age:  30,
					Comments: []Comment{
						{
							Text: "Great post!",
						},
						{
							Text: "I agree",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Test#2",
			data: []byte(`
				[ 
					{
						"name": "Jhon",
						"age": 30,
						"comments": [
							{"text": "Great post!"},
							{"text": "I agree"},
						]
					}
				]`),
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getUsersFromJSON(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUsersFromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUsersFromJSON() got = %v, want %v", got, tt.want)
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
	expected := "Name:  Jhon\nAge:  30\nComments: \n-  Great post!\n-  I agree\n\nName:  Alice\nAge:  25\nComments: \n-  Nice article\n\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
