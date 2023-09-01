package day9

import (
	"aoc2022/pkg/inputreader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Rope struct {
	Head        RopePiece
	Tail        RopePiece
	TailHistory map[string]struct{}
}

type RopePiece struct {
	X int
	Y int
}

// Move the head and track all locations the tail reaches
func (r *Rope) Move(dir string, amt int) error {
	xChange, yChange := getDelta(dir)
	for i := 0; i < amt; i++ {
		prevX := r.Head.X
		prevY := r.Head.Y

		r.Head.X = r.Head.X + xChange
		r.Head.Y = r.Head.Y + yChange

		// Move tail if head is moving more than 1 space away
		if math.Abs(float64(r.Head.X)-float64(r.Tail.X)) > 1 ||
			math.Abs(float64(r.Head.Y)-float64(r.Tail.Y)) > 1 {
			r.Tail.X = prevX
			r.Tail.Y = prevY
			r.TailHistory[fmt.Sprintf("%v,%v", r.Tail.X, r.Tail.Y)] = struct{}{}
		}

	}

	return nil
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

// Process all moves and return how many unique positions were visited by the tail
func ProcessMoves(lines []string) (int, error) {
	r := &Rope{TailHistory: map[string]struct{}{"0,0": {}}}

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

	ans, err := ProcessMoves(lines)
	if err != nil {
		return nil, err
	}

	return ans, nil
}

func Part2() (any, error) {

	return 0, nil
}
