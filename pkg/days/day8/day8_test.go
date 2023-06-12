package day8

import (
	"testing"
)

func Test_setVisible(t *testing.T) {
	tests := []struct {
		name        string
		arr         []*Tree
		wantVisible []bool
	}{
		{
			name: "[0, 0, 0]",
			arr: []*Tree{
				{Height: 0},
				{Height: 0},
				{Height: 0},
			},
			wantVisible: []bool{true, false, true},
		},
		{
			name: "All already set to visible",
			arr: []*Tree{
				{
					Height:  3,
					Visible: true,
				},
				{
					Height:  1,
					Visible: true,
				},
				{
					Height:  3,
					Visible: true,
				},
			},
			wantVisible: []bool{true, true, true},
		},
		{
			name: "[0, 1, 0, 0, 0, 1, 0]",
			arr: []*Tree{
				{Height: 0},
				{Height: 1},
				{Height: 0},
				{Height: 0},
				{Height: 0},
				{Height: 0},
				{Height: 1},
				{Height: 0},
			},
			wantVisible: []bool{true, true, false, false, false, false, true, true},
		},
		{
			name: "[0, 1, 2, 3, 4, 5, 6]",
			arr: []*Tree{
				{Height: 0},
				{Height: 1},
				{Height: 2},
				{Height: 3},
				{Height: 4},
				{Height: 5},
				{Height: 6},
			},
			wantVisible: []bool{true, true, true, true, true, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setVisible(tt.arr)
			for i, tree := range tt.arr {
				if tree.Visible != tt.wantVisible[i] {
					t.Errorf("buildForest() = %v, want %v", tree.Visible, tt.wantVisible)
				}
			}
		})
	}
}
