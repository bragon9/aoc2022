package day9

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// Return test rope assuming
//
//	coords ex: ["3 1", "3 0"]
//	history ex: ["3,1","299,2"]
func makeTestRope(coords []string, history []string) *Rope {
	var pieces []*RopePiece
	for i := 0; i < len(coords); i++ {
		fields := strings.Fields(coords[i])
		x, _ := strconv.Atoi(fields[0])
		y, _ := strconv.Atoi(fields[1])
		p := &RopePiece{
			X: x,
			Y: y,
		}
		pieces = append(pieces, p)
	}

	tailHistory := make(map[string]struct{}, 0)
	for i := 0; i < len(history); i++ {
		tailHistory[history[i]] = struct{}{}
	}

	return &Rope{
		Pieces:      pieces,
		TailHistory: tailHistory,
	}
}

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
			r: makeTestRope(
				[]string{
					"0 0",
					"0 0",
				},
				[]string{"0,0"},
			),
			args: args{
				dir: "R",
				amt: 4,
			},
			wantRope: makeTestRope(
				[]string{
					"4 0",
					"3 0",
				},
				[]string{
					"0,0",
					"1,0",
					"2,0",
					"3,0",
				},
			),
		},
		{
			name: "4,0 3,0 U 4",
			r: makeTestRope(
				[]string{
					"4 0",
					"3 0",
				},
				[]string{
					"3,0",
				},
			),
			args: args{
				dir: "U",
				amt: 4,
			},
			wantRope: makeTestRope(
				[]string{
					"4 4",
					"4 3",
				},
				[]string{
					"3,0",
					"4,1",
					"4,2",
					"4,3",
				},
			),
		},
		{
			name: "4,4 4,3 L 3",
			r: makeTestRope(
				[]string{
					"4 4",
					"4 3",
				},
				[]string{
					"4,3",
				},
			),
			args: args{
				dir: "L",
				amt: 3,
			},
			wantRope: makeTestRope(
				[]string{
					"1 4",
					"2 4",
				},
				[]string{
					"4,3",
					"3,4",
					"2,4",
				},
			),
		},
		{
			name: "1,4 2,4 D 1",
			r: makeTestRope(
				[]string{
					"1 4",
					"2 4",
				},
				[]string{
					"2,4",
				},
			),
			args: args{
				dir: "D",
				amt: 1,
			},
			wantRope: makeTestRope(
				[]string{
					"1 3",
					"2 4",
				},
				[]string{
					"2,4",
				},
			),
		},
		{
			name: "Head left of tail, move 2R, tail should not move",
			r: makeTestRope(
				[]string{
					"-1 0",
					"0 0",
				},
				[]string{
					"0,0",
				},
			),
			args: args{
				dir: "R",
				amt: 2,
			},
			wantRope: makeTestRope(
				[]string{
					"1 0",
					"0 0",
				},
				[]string{
					"0,0",
				},
			),
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

	r := makeTestRope(
		[]string{
			"0 0",
			"0 0",
		},
		[]string{"0,0"},
	)

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
