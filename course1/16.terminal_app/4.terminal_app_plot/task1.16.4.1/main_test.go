package main

import (
	"fmt"
	"testing"
)

func Test_requestAPI(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		wantErr bool
	}{
		{
			name:    "1",
			args:    "bitcoin",
			wantErr: false,
		},
		{
			name:    "2",
			args:    "litecoin",
			wantErr: false,
		},
		{
			name:    "3",
			args:    "ethereum",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := requestAPI(tt.args)
			fmt.Println(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("requestAPI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_requestAPI2(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		wantErr bool
	}{
		{
			name:    "1",
			args:    "BTC_USD",
			wantErr: false,
		},
		{
			name:    "2",
			args:    "LTC_USD",
			wantErr: false,
		},
		{
			name:    "3",
			args:    "ETH_USD",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := requestAPI2(tt.args)
			fmt.Println(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("requestAPI2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
