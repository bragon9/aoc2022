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
	r := MakeRope(len(coords))

	curr := r.Head
	for _, coord := range coords {
		fields := strings.Fields(coord)
		x, _ := strconv.Atoi(fields[0])
		y, _ := strconv.Atoi(fields[1])
		curr.X = x
		curr.Y = y
		curr = curr.Next
	}

	for _, coord := range history {
		r.Tail.History[coord] = struct{}{}
	}

	return r
}

func TestRope_MoveLength2(t *testing.T) {
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
				[]string{"0,0"},
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
				[]string{"0,0"},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Move(tt.args.dir, tt.args.amt); (err != nil) != tt.wantErr {
				t.Errorf("Rope.Move() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.wantRope.Tail.History, tt.r.Tail.History) {
				t.Errorf("unexpected rope travel")
			}
		})
	}
}

func TestRopeWeirdMovement(t *testing.T) {
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
			// ......
			// ......
			// ......
			// ....H.
			// 4321..  (4 covers 5, 6, 7, 8, 9, s)
			//
			// == U 1 ==
			// ......
			// ......
			// ....H.
			// .4321.
			// 5.....  (5 covers 6, 7, 8, 9, s)
			name: "length 10 U 8",
			r: makeTestRope(
				[]string{
					"4 1", // H
					"4 0", // 1
					"3 0", // 2
					"2 0", // 3
					"1 0", // 4
					"0 0", // 5
					"0 0",
					"0 0",
					"0 0",
					"0 0",
				},
				[]string{"0,0"},
			),
			args: args{
				dir: "U",
				amt: 8,
			},
			wantRope: makeTestRope(
				[]string{
					"5 8", // H
					"5 7", // 1
					"5 6", // 2
					"5 5", // 3
					"5 4", // 4
					"4 4", // 5
					"3 3", // 6
					"2 2", // 7
					"1 1", // 8
					"0 0", // 9
				},
				[]string{"0,0"},
			),
		},
		{
			// ...987654321H
			// .............
			// .............
			// .............
			// .............
			// .............
			// .............
			// .............
			// .............
			// .............
			// ............. (9 covers s)
			//
			// == D 10 ==
			// ...s.........
			// .............
			// .............
			// .............
			// .............
			// ........98765
			// ............4
			// ............3
			// ............2
			// ............1
			// ............H
			name: "horizontal line D 10",
			r: makeTestRope(
				[]string{
					"9 0", // H
					"8 0", // 1
					"7 0", // 2
					"6 0", // 3
					"5 0", // 4
					"4 0", // 5
					"3 0", // 6
					"2 0", // 7
					"1 0", // 8
					"0 0", // 9
				},
				[]string{"0,0"},
			),
			args: args{
				dir: "D",
				amt: 10,
			},
			wantRope: makeTestRope(
				[]string{

					"9 -10", // H
					"9 -9",  // 1
					"9 -8",  // 2
					"9 -7",  // 3
					"9 -6",  // 4
					"9 -5",  // 5
					"8 -5",  // 6
					"7 -5",  // 7
					"6 -5",  // 8
					"5 -5",  // 9
				},
				[]string{
					"0,0",
					"1,-1",
					"2,-2",
					"3,-3",
					"4,-4",
					"5,-5",
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Move(tt.args.dir, tt.args.amt); (err != nil) != tt.wantErr {
				t.Errorf("Rope.Move() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.wantRope.Tail.History, tt.r.Tail.History) {
				t.Errorf("unexpected rope travel")
			}
		})
	}
}

func TestRope_MoveLength10(t *testing.T) {
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
			name: "10 pieces, move R 4",
			r: makeTestRope(
				[]string{
					"0 0",
					"0 0",
					"0 0",
					"0 0",
					"0 0",
					"0 0",
					"0 0",
					"0 0",
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
					"4 0", // H
					"3 0", // 1
					"2 0", // 2
					"1 0", // 3
					"0 0",
					"0 0",
					"0 0",
					"0 0",
					"0 0",
					"0 0",
				},
				[]string{"0,0"},
			),
		},
		// ..........................
		// ..........................
		// ..........................
		// ...........54321H.........  (5 covers 6, 7, 8, 9, s)
		//
		// == U 8 ==
		// ................H.........
		// ................1.........
		// ................2.........
		// ................3.........
		// ...............54.........
		// ..............6...........
		// .............7............
		// ............8.............
		// ...........9..............  (9 covers s)
		//
		{
			name: "length 10 U 8",
			r: makeTestRope(
				[]string{
					"5 0", // H
					"4 0", // 1
					"3 0", // 2
					"2 0", // 3
					"1 0", // 4
					"0 0", // 5
					"0 0",
					"0 0",
					"0 0",
					"0 0",
				},
				[]string{"0,0"},
			),
			args: args{
				dir: "U",
				amt: 8,
			},
			wantRope: makeTestRope(
				[]string{
					"5 8", // H
					"5 7", // 1
					"5 6", // 2
					"5 5", // 3
					"5 4", // 4
					"4 4", // 5 - this is the broken one
					"3 3", // 6
					"2 2", // 7
					"1 1", // 8
					"0 0", // 9
				},
				[]string{"0,0"},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Move(tt.args.dir, tt.args.amt); (err != nil) != tt.wantErr {
				t.Errorf("Rope.Move() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.wantRope.Tail.History, tt.r.Tail.History) {
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

	if !reflect.DeepEqual(r.Tail.History, expectedHistory) {
		t.Errorf("unexpected history")
	}
}

func TestSampleInput_length10(t *testing.T) {
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
			"0 0",
			"0 0",
			"0 0",
			"0 0",
			"0 0",
			"0 0",
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
	}

	if !reflect.DeepEqual(r.Tail.History, expectedHistory) {
		t.Errorf("unexpected history")
	}
}

func TestPart2SampleInput(t *testing.T) {
	lines := []string{
		"R 5",
		"U 8",
		"L 8",
		"D 3",
		"R 17",
		"D 10",
		"L 25",
		"U 20",
	}

	r := makeTestRope(
		[]string{
			"0 0",
			"0 0",
			"0 0",
			"0 0",
			"0 0",
			"0 0",
			"0 0",
			"0 0",
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
		"0,0":   {},
		"1,1":   {},
		"2,2":   {},
		"1,3":   {},
		"2,4":   {},
		"3,5":   {},
		"4,5":   {},
		"5,5":   {},
		"6,4":   {},
		"7,3":   {},
		"8,2":   {},
		"9,1":   {},
		"10,0":  {},
		"9,-1":  {},
		"8,-2":  {},
		"7,-3":  {},
		"6,-4":  {},
		"5,-5":  {},
		"4,-5":  {},
		"3,-5":  {},
		"2,-5":  {},
		"1,-5":  {},
		"0,-5":  {},
		"-1,-5": {},
		"-2,-5": {},
		"-3,-4": {},
		"-4,-3": {},
		"-5,-2": {},
		"-6,-1": {},
		"-7,0":  {},
		"-8,1":  {},
		"-9,2":  {},
		"-10,3": {},
		"-11,4": {},
		"-11,5": {},
		"-11,6": {},
	}

	if !reflect.DeepEqual(r.Tail.History, expectedHistory) {
		t.Errorf("unexpected history")
	}
}
