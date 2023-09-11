package day10

import (
	"aoc2022/pkg/inputreader"
	"strconv"
	"strings"
)

const (
	NoOp  string = "noop"
	AddOp string = "addx"
)

type CRT struct {
	Cycle        int
	Value        int
	SignalChecks map[int]int
}

func (crt *CRT) NextCycle() {
	crt.Cycle += 1
	if _, ok := crt.SignalChecks[crt.Cycle]; ok {
		crt.SignalChecks[crt.Cycle] = crt.Value
	}
}

func (crt *CRT) AddOp(i int) {
	crt.NextCycle()
	crt.NextCycle()
	crt.Value += i
}

func (crt *CRT) NoOp() {
	crt.NextCycle()
}

func (crt *CRT) GetStrengthTotal() int {
	var total int
	for k, v := range crt.SignalChecks {
		total += (k * v)
	}

	return total
}

func (crt *CRT) ProcessLine(line string) error {
	fields := strings.Fields(line)
	op := fields[0]
	switch op {
	case NoOp:
		crt.NoOp()
	case AddOp:
		amt, err := strconv.Atoi(fields[1])
		if err != nil {
			return err
		}
		crt.AddOp(amt)
	}

	return nil
}

func Part1() (any, error) {
	lines, err := inputreader.ReadLines("inputs/day10/1.txt")
	if err != nil {
		return 0, err
	}

	crt := &CRT{
		Cycle: 0,
		Value: 1,
		SignalChecks: map[int]int{
			20:  0,
			60:  0,
			100: 0,
			140: 0,
			180: 0,
			220: 0,
		},
	}

	for _, line := range lines {
		err := crt.ProcessLine(line)
		if err != nil {
			return nil, err
		}
	}

	return crt.GetStrengthTotal(), nil
}

func Part2() (any, error) {
	// lines, err := inputreader.ReadLines("inputs/day10/1.txt")
	// if err != nil {
	// 	return 0, err
	// }

	return nil, nil
}
