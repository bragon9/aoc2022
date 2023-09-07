package day9

import (
	"aoc2022/pkg/inputreader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Rope struct {
	Pieces      []*RopePiece
	TailHistory map[string]struct{}
}

type RopePiece struct {
	X int
	Y int
}

// Make a rope of given size
func MakeRope(l int) *Rope {
	var pieces []*RopePiece

	for i := 0; i < l; i++ {
		pieces = append(pieces, &RopePiece{})
	}
	return &Rope{
		Pieces:      pieces,
		TailHistory: map[string]struct{}{"0,0": {}},
	}
}

func getDelta(dir string) (int, int) {
	switch dir {
	case "U":
		return 0, 1
	case "R":
		return 1, 0
	case "D":
		return 0, -1
	case "L":
		return -1, 0
	}

	return 0, 0
}

// Move the head and track all locations the tail reaches
func (r *Rope) Move(dir string, amt int) error {
	xChange, yChange := getDelta(dir)
	for i := 0; i < amt; i++ {
		// Loop through and move all pieces except for the last
		for j := 0; j < len(r.Pieces)-1; j++ {
			curr := r.Pieces[j]
			next := r.Pieces[j+1]

			// Save off values, we will move next piece to here potentially
			currX := curr.X
			currY := curr.Y

			curr.X = curr.X + xChange
			curr.Y = curr.Y + yChange

			// Update next piece if it is further than 1 space away
			if math.Abs(float64(curr.X)-float64(next.X)) > 1 ||
				math.Abs(float64(curr.Y)-float64(next.Y)) > 1 {
				next.X = currX
				next.Y = currY
			}

			// If tail was just moved, update history
			if j == len(r.Pieces)-2 {
				r.TailHistory[fmt.Sprintf("%v,%v", next.X, next.Y)] = struct{}{}
			}

		}

	}

	return nil
}

// Process all moves and return how many unique positions were visited by the tail
func (r *Rope) ProcessMoves(lines []string) (int, error) {
	for _, line := range lines {
		fields := strings.Fields(line)
		dir := fields[0]
		amt, err := strconv.Atoi(fields[1])
		if err != nil {
			return 0, err
		}

		err = r.Move(dir, amt)
		if err != nil {
			return 0, err
		}
	}

	return len(r.TailHistory), nil
}

func Part1() (any, error) {
	lines, err := inputreader.ReadLines("inputs/day9/1.txt")
	if err != nil {
		return 0, err
	}

	r := MakeRope(2)
	ans, err := r.ProcessMoves(lines)
	if err != nil {
		return nil, err
	}

	return ans, nil
}

func Part2() (any, error) {

	return 0, nil
}
