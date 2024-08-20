package main

import "testing"

func Test_getJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []User
		want    string
		wantErr bool
	}{
		{
			name: "Test#1",
			data: []User{
				{
					Name: "R2D2",
					Age:  18,
					Comments: []Comment{
						{
							Text: "Comment 1",
						},
						{
							Text: "Comment 2",
						},
						{
							Text: "Comment 3",
						},
					},
				},
				{
					Name: "Sergey",
					Age:  29,
					Comments: []Comment{
						{
							Text: "Dsd faf",
						},
						{
							Text: "Cozvmmdfgvcxbent adf",
						},
						{
							Text: "Cosdf fasdfa mment fsd",
						},
					},
				},
			},
			want:    "[{\"name\":\"R2D2\",\"age\":18,\"comments\":[{\"text\":\"Comment 1\"},{\"text\":\"Comment 2\"},{\"text\":\"Comment 3\"}]},{\"name\":\"Sergey\",\"age\":29,\"comments\":[{\"text\":\"Dsd faf\"},{\"text\":\"Cozvmmdfgvcxbent adf\"},{\"text\":\"Cosdf fasdfa mment fsd\"}]}]",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getJSON(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("getJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
