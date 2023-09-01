package day9

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestRope_Move(t *testing.T) {
	type args struct {
		dir string
		amt int
	}
	tests := []struct {
		name     string
		r        *Rope
		args     args
		wantErr  bool
		wantRope *Rope
	}{
		{
			name: "0,0 0,0 R 4",
			r: &Rope{
				Head: RopePiece{
					X: 0,
					Y: 0,
				},
				Tail: RopePiece{
					X: 0,
					Y: 0,
				},
				TailHistory: map[string]struct{}{"0,0": {}},
			},
			args: args{
				dir: "R",
				amt: 4,
			},
			wantRope: &Rope{
				Head: RopePiece{
					X: 4,
					Y: 0,
				},
				Tail: RopePiece{
					X: 3,
					Y: 0,
				},
				TailHistory: map[string]struct{}{
					"0,0": {},
					"1,0": {},
					"2,0": {},
					"3,0": {},
				},
			},
		},
		{
			name: "4,0 3,0 U 4",
			r: &Rope{
				Head: RopePiece{
					X: 4,
					Y: 0,
				},
				Tail: RopePiece{
					X: 3,
					Y: 0,
				},
				TailHistory: map[string]struct{}{"3,0": {}},
			},
			args: args{
				dir: "U",
				amt: 4,
			},
			wantRope: &Rope{
				Head: RopePiece{
					X: 4,
					Y: 4,
				},
				Tail: RopePiece{
					X: 4,
					Y: 3,
				},
				TailHistory: map[string]struct{}{
					"3,0": {},
					"4,1": {},
					"4,2": {},
					"4,3": {},
				},
			},
		},
		{
			name: "4,4 4,3 L 3",
			r: &Rope{
				Head: RopePiece{
					X: 4,
					Y: 4,
				},
				Tail: RopePiece{
					X: 4,
					Y: 3,
				},
				TailHistory: map[string]struct{}{"4,3": {}},
			},
			args: args{
				dir: "L",
				amt: 3,
			},
			wantRope: &Rope{
				Head: RopePiece{
					X: 1,
					Y: 4,
				},
				Tail: RopePiece{
					X: 2,
					Y: 4,
				},
				TailHistory: map[string]struct{}{
					"4,3": {},
					"3,4": {},
					"2,4": {},
				},
			},
		},
		{
			name: "1,4 2,4 D 1",
			r: &Rope{
				Head: RopePiece{
					X: 1,
					Y: 4,
				},
				Tail: RopePiece{
					X: 2,
					Y: 4,
				},
				TailHistory: map[string]struct{}{"2,4": {}},
			},
			args: args{
				dir: "D",
				amt: 1,
			},
			wantRope: &Rope{
				Head: RopePiece{
					X: 1,
					Y: 3,
				},
				Tail: RopePiece{
					X: 2,
					Y: 4,
				},
				TailHistory: map[string]struct{}{
					"2,4": {},
				},
			},
		},
		{
			name: "Head left of tail, move 2R, tail should not move",
			r: &Rope{
				Head: RopePiece{
					X: -1,
					Y: 0,
				},
				Tail: RopePiece{
					X: 0,
					Y: 0,
				},
				TailHistory: map[string]struct{}{"0,0": {}},
			},
			args: args{
				dir: "R",
				amt: 2,
			},
			wantRope: &Rope{
				Head: RopePiece{
					X: 1,
					Y: 0,
				},
				Tail: RopePiece{
					X: 0,
					Y: 0,
				},
				TailHistory: map[string]struct{}{
					"0,0": {},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Move(tt.args.dir, tt.args.amt); (err != nil) != tt.wantErr {
				t.Errorf("Rope.Move() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.wantRope, tt.r) {
				t.Errorf("unexpected rope travel")
			}
		})
	}
}

func TestSampleInput(t *testing.T) {
	lines := []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	}

	r := &Rope{TailHistory: map[string]struct{}{"0,0": {}}}

	for _, line := range lines {
		fields := strings.Fields(line)
		dir := fields[0]
		amt, err := strconv.Atoi(fields[1])
		if err != nil {
			t.Errorf("error converting")
		}

		err = r.Move(dir, amt)
		if err != nil {
			t.Errorf("error moving")
		}
	}

	expectedHistory := map[string]struct{}{
		"0,0": {},
		"1,0": {},
		"2,0": {},
		"3,0": {},
		"4,1": {},
		"1,2": {},
		"2,2": {},
		"3,2": {},
		"4,2": {},
		"3,3": {},
		"4,3": {},
		"2,4": {},
		"3,4": {},
	}

	if !reflect.DeepEqual(r.TailHistory, expectedHistory) {
		t.Errorf("unexpected history")
	}

}
