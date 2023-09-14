package day11

import (
	"reflect"
	"testing"
)

func Test_parseMonkeyNum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test 1",
			args: args{
				s: "Monkey 4:",
			},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseMonkeyNum(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMonkeyNum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseMonkeyNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseItems(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		wantItems []*Item
		wantErr   bool
	}{
		{
			name: "test 1",
			args: args{
				s: "Starting items: 54, 96, 82, 78, 69",
			},
			wantItems: []*Item{
				{54},
				{96},
				{82},
				{78},
				{69},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItems, err := parseItems(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItems, tt.wantItems) {
				t.Errorf("parseItems() = %v, want %v", gotItems, tt.wantItems)
			}
		})
	}
}

func Test_parseOperation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    func(i int) int
		wantErr bool
	}{
		{
			name: "old * 7",
			args: args{
				s: "new = old * 7",
			},
			want: func(i int) int {
				return i * 7
			},
			wantErr: false,
		},
		{
			name: "old + 5",
			args: args{
				s: "new = old + 5",
			},
			want: func(i int) int {
				return i + 5
			},
			wantErr: false,
		},
		{
			name: "old * old",
			args: args{
				s: "new = old * old",
			},
			want: func(i int) int {
				return i * i
			},
			wantErr: false,
		},
		{
			name: "old * 17",
			args: args{
				s: "new = old * 17",
			},
			want: func(i int) int {
				return i * 17
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseOperation(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseOperation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got(3) != tt.want(3) {
				t.Errorf("parseOperation() failed")
			}
		})
	}
}
