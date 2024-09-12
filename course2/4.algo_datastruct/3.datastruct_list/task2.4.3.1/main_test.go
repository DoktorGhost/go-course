package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {

	tests := []struct {
		name string
		args []Commit
		want []Commit
	}{
		{
			name: "1",
			args: []Commit{
				{Message: "aa", UUID: "adad", Date: "2006-01-02"},
				{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
				{Message: "aa", UUID: "asdfdad", Date: "2006-03-03"},
				{Message: "aa", UUID: "asadfdad", Date: "2008-01-05"},
				{Message: "aa", UUID: "adaasgdd", Date: "2001-01-02"},
			},
			want: []Commit{
				{Message: "aa", UUID: "adaasgdd", Date: "2001-01-02"},
				{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
				{Message: "aa", UUID: "adad", Date: "2006-01-02"},
				{Message: "aa", UUID: "asdfdad", Date: "2006-03-03"},
				{Message: "aa", UUID: "asadfdad", Date: "2008-01-05"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if QuickSort(tt.args); !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("selectionSort() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Init(t *testing.T) {
	type want struct {
		len     int
		message string
		UUID    string
		date    string
		idx     string
	}
	tests := []struct {
		name string
		args []Commit
		want want
	}{
		{
			name: "head",
			args: []Commit{
				Commit{Message: "a", UUID: "a11", Date: "2006-01-02"},
				Commit{Message: "b", UUID: "b12", Date: "2005-01-02"},
				Commit{Message: "c", UUID: "c13", Date: "2004-01-02"},
				Commit{Message: "d", UUID: "d14", Date: "2008-01-02"},
				Commit{Message: "e", UUID: "e15", Date: "2009-01-02"},
			},
			want: want{5, "a", "a11", "2006-01-02", "head"},
		},
		{
			name: "tail",
			args: []Commit{
				Commit{Message: "a", UUID: "a11", Date: "2006-01-02"},
				Commit{Message: "b", UUID: "b12", Date: "2005-01-02"},
				Commit{Message: "c", UUID: "c13", Date: "2004-01-02"},
				Commit{Message: "d", UUID: "d14", Date: "2008-01-02"},
			},
			want: want{4, "d", "d14", "2008-01-02", "tail"},
		},
		{
			name: "1 element head",
			args: []Commit{Commit{Message: "a", UUID: "a11", Date: "2006-01-02"}},
			want: want{1, "a", "a11", "2006-01-02", "head"},
		},
		{
			name: "1 element tail",
			args: []Commit{Commit{Message: "a", UUID: "a11", Date: "2006-01-02"}},
			want: want{1, "a", "a11", "2006-01-02", "tail"},
		},
		{
			name: "nil",
			args: []Commit{},
			want: want{0, "", "", "", "head"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{}
			d.Init(tt.args)
			node := d.head
			if tt.want.idx == "head" {
				node = d.head
			} else if tt.want.idx == "tail" {
				node = d.tail
			}
			if tt.want.len == 0 {
				if !reflect.DeepEqual(d.len, tt.want.len) {
					t.Errorf("len() = %v, want %v", d.len, tt.want.len)
				}
			} else {

				if !reflect.DeepEqual(d.len, tt.want.len) {
					t.Errorf("len() = %v, want %v", d.len, tt.want.len)
				}
				if !reflect.DeepEqual(node.data.UUID, tt.want.UUID) {
					t.Errorf("Message() = %v, want %v", node.data.UUID, tt.want.UUID)
				}
				if !reflect.DeepEqual(node.data.Date, tt.want.date) {
					t.Errorf("Message() = %v, want %v", node.data.Date, tt.want.date)
				}
				if !reflect.DeepEqual(node.data.Message, tt.want.message) {
					t.Errorf("Message() = %v, want %v", node.data.Message, tt.want.message)
				}
			}
		})
	}
}

func TestDoubleLinkedList_LoadData(t *testing.T) {
	type want struct {
		len     int
		message string
		UUID    string
		date    string
		idx     string
	}

	tests := []struct {
		name string
		path string
		want want
	}{
		{
			name: "valid file head",
			path: "testfile.json",
			want: want{5, "c", "c13", "2004-01-02", "head"},
		},

		{
			name: "valid file tail",
			path: "testfile.json",
			want: want{5, "e", "e15", "2009-01-02", "tail"},
		},

		{
			name: "nil path",
			path: "",
			want: want{},
		},
		{
			name: "bad path",
			path: "asdaxcv",
			want: want{},
		},
		{
			name: "bad json",
			path: "testfilebad.json",
			want: want{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{}

			if tt.want.len == 0 {
				if err := d.LoadData(tt.path); err == nil {
					t.Errorf("LoadData() error = %v, wantErr %v", err, nil)
				}
			} else {
				if err := d.LoadData(tt.path); err != nil {
					t.Errorf("LoadData() error = %v, wantErr %v", err, nil)
				}

				node := d.head
				if tt.want.idx == "head" {
					node = d.head
				} else if tt.want.idx == "tail" {
					node = d.tail
				}

				if !reflect.DeepEqual(d.len, tt.want.len) {
					t.Errorf("len() = %v, want %v", d.len, tt.want.len)
				}
				if !reflect.DeepEqual(node.data.UUID, tt.want.UUID) {
					t.Errorf("Message() = %v, want %v", node.data.UUID, tt.want.UUID)
				}
				if !reflect.DeepEqual(node.data.Date, tt.want.date) {
					t.Errorf("Message() = %v, want %v", node.data.Date, tt.want.date)
				}
				if !reflect.DeepEqual(node.data.Message, tt.want.message) {
					t.Errorf("Message() = %v, want %v", node.data.Message, tt.want.message)
				}
			}

		})
	}
}

func TestDoubleLinkedList_Len(t *testing.T) {
	tests := []struct {
		name string
		args []Commit
		want int
	}{
		{
			name: "normal",
			args: []Commit{
				{Message: "aa", UUID: "adad", Date: "2006-01-02"},
				{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
				{Message: "aa", UUID: "asdfdad", Date: "2006-03-03"},
				{Message: "aa", UUID: "asadfdad", Date: "2008-01-05"},
				{Message: "aa", UUID: "adaasgdd", Date: "2001-01-02"},
			},
			want: 5,
		},
		{
			name: "0 element",
			args: []Commit{},
			want: 0,
		},
		{
			name: "normal2",
			args: []Commit{
				{Message: "aa", UUID: "adad", Date: "2006-01-02"},
				{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
				{Message: "aa", UUID: "asdfdad", Date: "2006-03-03"},
				{Message: "aa", UUID: "asadfdad", Date: "2008-01-05"},
				{Message: "aa", UUID: "adaasgdd", Date: "2001-01-02"},
				{Message: "aa", UUID: "adad", Date: "2006-01-02"},
				{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
				{Message: "aa", UUID: "asdfdad", Date: "2006-03-03"},
				{Message: "aa", UUID: "asadfdad", Date: "2008-01-05"},
				{Message: "aa", UUID: "adaasgdd", Date: "2001-01-02"},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{}
			d.Init(tt.args)

			if got := d.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_SetCurrent(t *testing.T) {
	type args struct {
		n int
		c []Commit
	}
	tests := []struct {
		name     string
		args     args
		currUUID string
	}{
		{
			name: "normal1",
			args: args{
				n: 2,
				c: []Commit{
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
					{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
					{Message: "aa", UUID: "adaasfasgdd", Date: "2001-01-02"},
				},
			},
			currUUID: "asfsdfdad",
		},
		{
			name: "normal2",
			args: args{
				n: 1,
				c: []Commit{
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
					{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
					{Message: "aa", UUID: "adaasfasgdd", Date: "2001-01-02"},
				},
			},
			currUUID: "adadasd",
		},
		{
			name: "normal3",
			args: args{
				n: 4,
				c: []Commit{
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
					{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
					{Message: "aa", UUID: "adaasfasgdd", Date: "2001-01-02"},
				},
			},
			currUUID: "adaasfasgdd",
		},
		{
			name: "normal4",
			args: args{
				n: 3,
				c: []Commit{
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
					{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
					{Message: "aa", UUID: "adaasfasgdd", Date: "2001-01-02"},
				},
			},
			currUUID: "asadfsdfgdad",
		},

		{
			name: "nil",
			args: args{
				n: 3,
				c: []Commit{},
			},
			currUUID: "",
		},
		{
			name: "bad",
			args: args{
				n: 5,
				c: []Commit{
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
				},
			},
			currUUID: "",
		},
		{
			name: "bad",
			args: args{
				n: -5,
				c: []Commit{
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
				},
			},
			currUUID: "",
		},
		{
			name: "normal 0",
			args: args{
				n: 0,
				c: []Commit{
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
				},
			},
			currUUID: "adsfdad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{}
			d.Init(tt.args.c)
			if tt.name == "nil" || tt.name == "bad" {
				if err := d.SetCurrent(tt.args.n); err == nil {
					t.Errorf("SetCurrent() error = %v, wantErr %s", err, "not nil")
				}
			} else {

				if err := d.SetCurrent(tt.args.n); err != nil || d.curr.data.UUID != tt.currUUID {
					t.Errorf("SetCurrent() curr UUID = %v, want UUID %v", d.curr.data.UUID, tt.currUUID)
				}
			}
		})
	}
}

func TestDoubleLinkedList_Current(t *testing.T) {

	tests := []struct {
		name     string
		args     []Commit
		wantUUID string
	}{
		{
			name: "normal",
			args: []Commit{
				{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
				{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
				{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
				{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
				{Message: "aa", UUID: "adaasfasgdd", Date: "2001-01-02"},
			},
			wantUUID: "adaasfasgdd",
		},
		{
			name: "normal2",
			args: []Commit{
				{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
			},
			wantUUID: "adsfdad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{curr: nil}
			d.Init(tt.args)

			if got := d.Current(); !reflect.DeepEqual(got.data.UUID, tt.wantUUID) {
				t.Errorf("Current() = %v, want %v", got.data.UUID, tt.wantUUID)
			}
		})
	}
}

func TestDoubleLinkedList_Next(t *testing.T) {

	tests := []struct {
		name     string
		args     []Commit
		wantUUID string
	}{
		{
			name: "normal",
			args: []Commit{
				{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
				{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
				{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
				{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
				{Message: "aa", UUID: "adaasfasgdd", Date: "2001-01-02"},
			},
			wantUUID: "adaasfasgdd",
		},
		{
			name: "normal2",
			args: []Commit{
				{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
				{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
				{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
				{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
				{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
				{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
				{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
				{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
				{Message: "aa", UUID: "adaasfasgdd", Date: "2001-01-02"},
			},
			wantUUID: "adsfdad",
		},
		{
			name: "bad",
			args: []Commit{
				{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
				{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
				{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
			},
			wantUUID: "adsfdad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{}
			d.Init(tt.args)
			_ = d.SetCurrent(3)
			if tt.name == "bad" {
				if got := d.Next(); got != nil {
					t.Errorf("NextUUID = %v, want %v", got, nil)
				}
			} else {

				if got := d.Next(); !reflect.DeepEqual(got.data.UUID, tt.wantUUID) {
					t.Errorf("NextUUID = %v, want %v", got.data.UUID, tt.wantUUID)
				}
			}
		})
	}
}

func TestDoubleLinkedList_Prev(t *testing.T) {

	tests := []struct {
		name     string
		args     []Commit
		wantUUID string
	}{
		{
			name: "normal",
			args: []Commit{
				{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
				{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
				{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
				{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
				{Message: "aa", UUID: "adaasfasgdd", Date: "2001-01-02"},
			},
			wantUUID: "asfsdfdad",
		},
		{
			name: "normal2",
			args: []Commit{

				{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
				{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
				{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
				{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
				{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
				{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
				{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
				{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
				{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
				{Message: "aa", UUID: "adaasfasgdd", Date: "2001-01-02"},
			},
			wantUUID: "asfsdfdad",
		},
		{
			name: "bad",
			args: []Commit{

				{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
				{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
				{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
			},
			wantUUID: "asfsdfdad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{}
			d.Init(tt.args)
			if tt.name == "bad" {
				_ = d.SetCurrent(0)
				if got := d.Prev(); got != nil {
					t.Errorf("Prev UUID = %v, want %v", got, nil)
				}
			} else {

				_ = d.SetCurrent(3)
				if got := d.Prev(); !reflect.DeepEqual(got.data.UUID, tt.wantUUID) {
					t.Errorf("Prev UUID = %v, want %v", got.data.UUID, tt.wantUUID)
				}
			}
		})
	}
}

func TestDoubleLinkedList_Insert(t *testing.T) {
	type args struct {
		n       int
		commit  Commit
		commits []Commit
	}
	type want struct {
		prevUUID string
		nextUUID string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "bad",
			args: args{
				n:      -5,
				commit: Commit{Message: "test", UUID: "testUUID", Date: "2024-09-12"},
				commits: []Commit{
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
				},
			},
			want: want{prevUUID: "adsfdad", nextUUID: "asfsdfdad"},
		},
		{
			name: "bad",
			args: args{
				n:      6,
				commit: Commit{Message: "test", UUID: "testUUID", Date: "2024-09-12"},
				commits: []Commit{
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
				},
			},
			want: want{prevUUID: "adsfdad", nextUUID: "asfsdfdad"},
		},
		{
			name: "normal",
			args: args{
				n:      5,
				commit: Commit{Message: "test", UUID: "testUUID", Date: "2024-09-12"},
				commits: []Commit{
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},

					//{Message: "test", UUID: "testUUID", Date: "2024-09-12"}
				},
			},
			want: want{prevUUID: "asadfsdfgdad"},
		},
		{
			name: "normal",
			args: args{
				n:      3,
				commit: Commit{Message: "test", UUID: "testUUID", Date: "2024-09-12"},
				commits: []Commit{
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
				},
			},
			want: want{prevUUID: "asfsdfdad", nextUUID: "adsfdad"},
		},
		{
			name: "normal",
			args: args{
				n:      0,
				commit: Commit{Message: "test", UUID: "testUUID", Date: "2024-09-12"},
				commits: []Commit{
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
				},
			},
			want: want{nextUUID: "adadasd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{}
			d.Init(tt.args.commits)
			err := d.Insert(tt.args.n, tt.args.commit)
			_ = d.SetCurrent(tt.args.n)

			if tt.name == "bad" {
				if err == nil {
					t.Errorf("Insert() got = %v, want = not nil", err)
				}
			} else {
				prev := d.curr.prev
				next := d.curr.next

				if tt.args.n == 0 {
					if !reflect.DeepEqual(next.data.UUID, tt.want.nextUUID) {
						t.Errorf("Next UUID = %v, want next %v", next.data.UUID, tt.want.nextUUID)
					}
				} else if tt.args.n == d.len-1 {
					if !reflect.DeepEqual(prev.data.UUID, tt.want.prevUUID) {
						t.Errorf("Prev UUID = %v, want prev %v", prev.data.UUID, tt.want.prevUUID)
					}
				} else {

					if !reflect.DeepEqual(prev.data.UUID, tt.want.prevUUID) {
						t.Errorf("Prev UUID = %v, want prev %v", prev.data.UUID, tt.want.prevUUID)
					}

					if !reflect.DeepEqual(next.data.UUID, tt.want.nextUUID) {
						t.Errorf("Next UUID = %v, want next %v", next.data.UUID, tt.want.nextUUID)
					}
				}

			}

		})
	}
}

func TestDoubleLinkedList_Push(t *testing.T) {

	type args struct {
		c       Commit
		commits []Commit
	}
	type want struct {
		prevUUID string
		len      int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "normal",
			args: args{
				c: Commit{Message: "test", UUID: "testUUID", Date: "2024-09-12"},
				commits: []Commit{
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
				},
			},
			want: want{
				prevUUID: "asadfsdfgdad",
				len:      6,
			},
		},
		{
			name: "bad",
			args: args{
				c: Commit{Message: "test", UUID: "testUUID", Date: "2024-09-12"},
				commits: []Commit{
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
				},
			},
			want: want{
				prevUUID: "asadfsdfgdad",
				len:      6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{}
			d.Init(tt.args.commits)

			if tt.name == "bad" {
				d.tail = nil
				err := d.Push(tt.args.c)
				if err == nil {
					t.Errorf("Push() got = %v, want = not nil", err)
				}
			} else {
				err := d.Push(tt.args.c)
				if err != nil {
					t.Errorf("Push() error = %v, wantErr nil", err)
				}
				if !reflect.DeepEqual(d.len, tt.want.len) {
					t.Errorf("Len = %v, want len %v", d.len, tt.want.len)
				}
				if !reflect.DeepEqual(d.tail.prev.data.UUID, tt.want.prevUUID) {
					t.Errorf("Prev UUID = %v, want prev %v", d.tail.data.UUID, tt.want.prevUUID)
				}
			}

		})
	}
}

func TestDoubleLinkedList_Delete(t *testing.T) {

	type args struct {
		n int
		c []Commit
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
	}{
		{
			name: "normal",
			args: args{
				n: 0,
				c: []Commit{
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
				},
			},
			wantLen: 4,
		},
		{
			name: "normal",
			args: args{
				n: 4,
				c: []Commit{
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
				},
			},
			wantLen: 4,
		},
		{
			name: "normal",
			args: args{
				n: 2,
				c: []Commit{
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
					{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
				},
			},
			wantLen: 4,
		},
		{
			name: "bad",
			args: args{
				n: -1,
				c: []Commit{
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
				},
			},
		},
		{
			name: "bad",
			args: args{
				n: 4,
				c: []Commit{
					{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
					{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{}
			d.Init(tt.args.c)
			err := d.Delete(tt.args.n)

			if tt.name == "bad" {
				if err == nil {
					t.Errorf("Delete() got = %v, want = not nil", err)
				}
			} else {
				if err != nil {
					t.Errorf("Delete() error = %v, wantErr nil", err)
				}
				if !reflect.DeepEqual(d.len, tt.wantLen) {
					t.Errorf("Delete() len = %v, wantLen %v", d.len, tt.wantLen)
				}
			}
		})
	}
}

func TestDoubleLinkedList_DeleteCurrent(t *testing.T) {
	tests := []struct {
		name    string
		args    []Commit
		wantLen int
	}{
		{
			name: "normal",
			args: []Commit{
				{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
				{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
				{Message: "aa", UUID: "asfsdfdad", Date: "2006-03-03"},
				{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
				{Message: "aa", UUID: "asadfsdfgdad", Date: "2008-01-05"},
			},

			wantLen: 4,
		},
		{
			name: "bad",
			args: []Commit{
				{Message: "sdaa", UUID: "adadasd", Date: "2001-05-02"},
				{Message: "aa", UUID: "adsfdad", Date: "2006-01-02"},
			},
		},
		{
			name: "bad",
			args: []Commit{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{}
			d.Init(tt.args)

			if tt.name == "bad" {
				d.curr = nil
				err := d.DeleteCurrent()
				if err == nil {
					t.Errorf("DeleteCurrent() got = %v, want = not nil", err)
				}
			} else {
				err := d.DeleteCurrent()
				if err != nil {
					t.Errorf("DeleteCurrent() error = %v, wantErr nill", err)
				}
				if !reflect.DeepEqual(d.len, tt.wantLen) {
					t.Errorf("DeleteCurrent() len = %v, wantLen %v", d.len, tt.wantLen)
				}
				_ = d.SetCurrent(0)
				if err != nil {
					t.Errorf("DeleteCurrent() error = %v, wantErr nill", err)
				}
				if !reflect.DeepEqual(d.len, tt.wantLen) {
					t.Errorf("DeleteCurrent() len = %v, wantLen %v", d.len, tt.wantLen-1)
				}
				_ = d.SetCurrent(d.len - 1)
				if err != nil {
					t.Errorf("DeleteCurrent() error = %v, wantErr nill", err)
				}
				if !reflect.DeepEqual(d.len, tt.wantLen) {
					t.Errorf("DeleteCurrent() len = %v, wantLen %v", d.len, tt.wantLen-1)
				}

			}
		})
	}
}
