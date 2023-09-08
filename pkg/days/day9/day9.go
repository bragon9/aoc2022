package day9

import (
	"aoc2022/pkg/inputreader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Linked list constructed of `*RopePiece`
type Rope struct {
	Head   *RopePiece
	Tail   *RopePiece
	Length int
}

// Representation of a "node"
type RopePiece struct {
	Number  int
	X       int
	Y       int
	Next    *RopePiece
	History map[string]struct{}
}

// Make a rope of given size
func MakeRope(l int) *Rope {
	head := &RopePiece{}
	next := head
	var tail *RopePiece
	for i := 0; i < l-1; i++ {
		next.Number = i
		next.Next = &RopePiece{}
		next = next.Next
		tail = next
	}
	tail.Number = l - 1
	tail.History = map[string]struct{}{"0,0": {}}
	return &Rope{
		Head:   head,
		Tail:   tail,
		Length: l,
	}
}

// Return the amount to increment X and Y from the given direction
func getDeltaFromDirection(dir string) (int, int) {
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

func getDeltaForNext(rp *RopePiece) (int, int) {
	var xChange int
	var yChange int
	switch {
	case rp.X > rp.Next.X:
		xChange = 1
	case rp.X < rp.Next.X:
		xChange = -1
	}
	switch {
	case rp.Y > rp.Next.Y:
		yChange = 1
	case rp.Y < rp.Next.Y:
		yChange = -1
	}

	return xChange, yChange
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

	return len(r.Tail.History), nil
}

// Handles moving the Head as directed in the input
func (r *Rope) Move(dir string, amt int) error {
	xChange, yChange := getDeltaFromDirection(dir)
	for i := 0; i < amt; i++ {
		r.Head.move(xChange, yChange)
	}
	return nil
}

// Move a piece of the rope
func (rp *RopePiece) move(xChange, yChange int) error {
	// Move the piece
	rp.X = rp.X + xChange
	rp.Y = rp.Y + yChange

	// See if the next section needs to move
	if rp.Next != nil {
		// Recursively move the rest of the rope, if necessary
		if math.Abs(float64(rp.X)-float64(rp.Next.X)) > 1 ||
			math.Abs(float64(rp.Y)-float64(rp.Next.Y)) > 1 {
			rp.Next.move(getDeltaForNext(rp))
		}
	} else {
		// Last piece -- update the tail history
		rp.History[fmt.Sprintf("%v,%v", rp.X, rp.Y)] = struct{}{}
	}

	return nil
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
	lines, err := inputreader.ReadLines("inputs/day9/1.txt")
	if err != nil {
		return 0, err
	}

	r := MakeRope(10)
	ans, err := r.ProcessMoves(lines)
	if err != nil {
		return nil, err
	}

	return ans, nil
}
