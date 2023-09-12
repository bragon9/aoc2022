package day10

import (
	"aoc2022/pkg/inputreader"
	"fmt"
	"math"
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
	Screen       [][]string
}

func (crt *CRT) PrintScreen() {
	for _, line := range crt.Screen {
		fmt.Println(line)
	}
}

func (crt *CRT) DrawPixel() {
	if crt.Screen == nil {
		return
	}
	pixel := (crt.Cycle - 1) % 40
	screen := crt.Cycle / 40
	if math.Abs(float64(crt.Value)-float64(pixel)) <= 1 {
		crt.Screen[screen][pixel] = "#"
	}
}

func (crt *CRT) NextCycle() {
	crt.Cycle += 1
	if _, ok := crt.SignalChecks[crt.Cycle]; ok {
		crt.SignalChecks[crt.Cycle] = crt.Value
	}
	crt.DrawPixel()
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
	lines, err := inputreader.ReadLines("inputs/day10/1.txt")
	if err != nil {
		return 0, err
	}

	crt := &CRT{
		Cycle: 0,
		Value: 1,
		Screen: [][]string{
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		},
	}

	for _, line := range lines {
		err := crt.ProcessLine(line)
		if err != nil {
			return nil, err
		}
	}

	crt.PrintScreen()

	return nil, nil
}
