package main

import (
	"bytes"
	"os"
	"testing"
)

func Test_filterSentence(t *testing.T) {
	type args struct {
		sentence string
		filter   map[string]bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test#1: Normal",
			args: args{
				sentence: "Go in my very good fathers hous from mam or not in",
				filter:   map[string]bool{"in": true, "good": true, "fathers": true},
			},
			want: "Go my very hous from mam or not",
		},
		{
			name: "Test#2: nil",
			args: args{
				sentence: "",
				filter:   map[string]bool{"in": true, "good": true, "fathers": true},
			},
			want: "",
		},
		{
			name: "Test#3: Normal",
			args: args{
				sentence: "in in",
				filter:   map[string]bool{"in": true, "good": true, "fathers": true},
			},
			want: "",
		},
		{
			name: "Test#4: Normal",
			args: args{
				sentence: " ",
				filter:   map[string]bool{"in": true, "good": true, "fathers": true},
			},
			want: "",
		},
		{
			name: "Test#5: Normal",
			args: args{
				sentence: "in     in   yes",
				filter:   map[string]bool{"in": true, "good": true, "fathers": true},
			},
			want: "yes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterSentence(tt.args.sentence, tt.args.filter); got != tt.want {
				t.Errorf("filterSentence() = %v, want %v", got, tt.want)
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
	expected := "Lorem dolor sit amet consectetur adipiscing\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
