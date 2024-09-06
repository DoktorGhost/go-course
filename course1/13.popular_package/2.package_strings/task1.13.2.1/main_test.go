package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestCountWordsInText(t *testing.T) {
	type args struct {
		txt   string
		words []string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "test#1",
			args: args{
				txt:   "Всем привет! На связи или не на связи ваш покорный слуга! Привет! И еще раз привет!",
				words: []string{"привет", "связи"},
			},
			want: map[string]int{"привет": 3, "связи": 2},
		},
		{
			name: "test#2",
			args: args{
				txt:   "",
				words: []string{"привет", "связи"},
			},
			want: map[string]int{"привет": 0, "связи": 0},
		},
		{
			name: "test#3",
			args: args{
				txt:   "привет связи",
				words: []string{"привет", "связи"},
			},
			want: map[string]int{"привет": 1, "связи": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountWordsInText(tt.args.txt, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountWordsInText() = %v, want %v", got, tt.want)
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
	expected := "map[amet:3 lorem:1 sit:3]\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
