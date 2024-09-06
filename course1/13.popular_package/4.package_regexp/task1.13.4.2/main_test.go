package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func Test_censorAds(t *testing.T) {
	type args struct {
		ads    []Ad
		censor map[string]string
	}
	tests := []struct {
		name string
		args args
		want []Ad
	}{
		{
			name: "Test#1",
			args: args{
				ads: []Ad{
					{
						Title:       "Куплю велосипед MeRiDa",
						Description: "Куплю велосипед meriDA в хорошем состоянии",
					},
				},
				censor: map[string]string{
					"велосипед merida": "телефон Apple",
				},
			},
			want: []Ad{

				{
					Title:       "Куплю телефон Apple",
					Description: "Куплю телефон Apple в хорошем состоянии",
				},
			},
		},
		{
			name: "Test#2",
			args: args{
				ads: []Ad{
					{
						Title:       "",
						Description: "",
					},
				},
				censor: map[string]string{
					"велосипед merida": "телефон Apple",
				},
			},
			want: nil,
		},
		{
			name: "Test#3",
			args: args{
				ads: []Ad{
					{
						Title:       "Куплю велосипед MeRiDa",
						Description: "Куплю велосипед meriDA в хорошем состоянии",
					},
				},
				censor: map[string]string{},
			},
			want: []Ad{
				{
					Title:       "Куплю велосипед MeRiDa",
					Description: "Куплю велосипед meriDA в хорошем состоянии",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := censorAds(tt.args.ads, tt.args.censor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("censorAds() = %v, want %v", got, tt.want)
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
	expected := 350
	if len(stdout.String()) != expected {
		t.Errorf("got %d, want %d", len(stdout.String()), expected)
	}
}
