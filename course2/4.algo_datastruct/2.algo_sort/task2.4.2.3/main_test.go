package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		arr1 []User
		arr2 []User
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{
			name: "Normal merge",
			args: args{
				arr1: []User{
					{ID: 1, Name: "Oleg", Age: 30},
					{ID: 5, Name: "FGd", Age: 15},
					{ID: 8, Name: "OfdlFDeg", Age: 22},
					{ID: 9, Name: "Olxvcbexg", Age: 15},
					{ID: 13, Name: "xvbOxvbleg", Age: 80},
				},
				arr2: []User{
					{ID: 3, Name: "Oleg", Age: 30},
					{ID: 7, Name: "FGd", Age: 15},
					{ID: 14, Name: "OfdlFDeg", Age: 22},
					{ID: 15, Name: "Olxvcbexg", Age: 15},
					{ID: 16, Name: "xvbOxvbleg", Age: 80},
				},
			},
			want: []User{
				{ID: 1, Name: "Oleg", Age: 30},
				{ID: 3, Name: "Oleg", Age: 30},
				{ID: 5, Name: "FGd", Age: 15},
				{ID: 7, Name: "FGd", Age: 15},
				{ID: 8, Name: "OfdlFDeg", Age: 22},
				{ID: 9, Name: "Olxvcbexg", Age: 15},
				{ID: 13, Name: "xvbOxvbleg", Age: 80},
				{ID: 14, Name: "OfdlFDeg", Age: 22},
				{ID: 15, Name: "Olxvcbexg", Age: 15},
				{ID: 16, Name: "xvbOxvbleg", Age: 80},
			},
		},
		{
			name: "args2 nil",
			args: args{
				arr1: []User{
					{ID: 1, Name: "Oleg", Age: 30},
					{ID: 5, Name: "FGd", Age: 15},
					{ID: 8, Name: "OfdlFDeg", Age: 22},
					{ID: 9, Name: "Olxvcbexg", Age: 15},
					{ID: 13, Name: "xvbOxvbleg", Age: 80},
				},
				arr2: []User{},
			},
			want: []User{
				{ID: 1, Name: "Oleg", Age: 30},
				{ID: 5, Name: "FGd", Age: 15},
				{ID: 8, Name: "OfdlFDeg", Age: 22},
				{ID: 9, Name: "Olxvcbexg", Age: 15},
				{ID: 13, Name: "xvbOxvbleg", Age: 80},
			},
		},
		{
			name: "args1 nil",
			args: args{
				arr1: []User{},
				arr2: []User{
					{ID: 3, Name: "Oleg", Age: 30},
					{ID: 7, Name: "FGd", Age: 15},
					{ID: 14, Name: "OfdlFDeg", Age: 22},
					{ID: 15, Name: "Olxvcbexg", Age: 15},
					{ID: 16, Name: "xvbOxvbleg", Age: 80},
				},
			},
			want: []User{
				{ID: 3, Name: "Oleg", Age: 30},
				{ID: 7, Name: "FGd", Age: 15},
				{ID: 14, Name: "OfdlFDeg", Age: 22},
				{ID: 15, Name: "Olxvcbexg", Age: 15},
				{ID: 16, Name: "xvbOxvbleg", Age: 80},
			},
		},
		{
			name: "arr1 < arr2",
			args: args{
				arr1: []User{
					{ID: 1, Name: "Oleg", Age: 30},
				},
				arr2: []User{
					{ID: 3, Name: "Oleg", Age: 30},
					{ID: 7, Name: "FGd", Age: 15},
					{ID: 14, Name: "OfdlFDeg", Age: 22},
					{ID: 15, Name: "Olxvcbexg", Age: 15},
					{ID: 16, Name: "xvbOxvbleg", Age: 80},
				},
			},
			want: []User{
				{ID: 1, Name: "Oleg", Age: 30},
				{ID: 3, Name: "Oleg", Age: 30},
				{ID: 7, Name: "FGd", Age: 15},
				{ID: 14, Name: "OfdlFDeg", Age: 22},
				{ID: 15, Name: "Olxvcbexg", Age: 15},
				{ID: 16, Name: "xvbOxvbleg", Age: 80},
			},
		},
		{
			name: "arr1 > arr2",
			args: args{
				arr1: []User{
					{ID: 3, Name: "Oleg", Age: 30},
					{ID: 7, Name: "FGd", Age: 15},
					{ID: 14, Name: "OfdlFDeg", Age: 22},
					{ID: 15, Name: "Olxvcbexg", Age: 15},
					{ID: 16, Name: "xvbOxvbleg", Age: 80},
				},
				arr2: []User{
					{ID: 1, Name: "Oleg", Age: 30},
				},
			},
			want: []User{
				{ID: 1, Name: "Oleg", Age: 30},
				{ID: 3, Name: "Oleg", Age: 30},
				{ID: 7, Name: "FGd", Age: 15},
				{ID: 14, Name: "OfdlFDeg", Age: 22},
				{ID: 15, Name: "Olxvcbexg", Age: 15},
				{ID: 16, Name: "xvbOxvbleg", Age: 80},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
