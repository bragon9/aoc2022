package day11

import (
	"aoc2022/pkg/inputreader"
	"fmt"
	"strconv"
	"strings"
)

type Troop struct {
	Monkeys [8]*Monkey
}

type Monkey struct {
	Items     []*Item
	Operation func(int) int
	Test      func(int) *Monkey
}
type Item struct {
	WorryLevel int
}

// All monkeys in the troop rummage in order
func (t Troop) PlayRound() error {
	for _, m := range t.Monkeys {
		err := m.Rummage()
		if err != nil {
			return err
		}
	}

	return nil
}

// Look through all items in order and throw them
func (m *Monkey) Rummage() error {
	for _, item := range m.Items {
		item.WorryLevel = m.Operation(item.WorryLevel) / 3
		// Throw to monkey determined by the test function
		m2 := m.Test(item.WorryLevel)
		m2.Items = append(m2.Items, item)
	}
	m.Items = []*Item{}

	return nil
}

func (t *Troop) parseTest(lines []string) (func(int) *Monkey, error) {
	d, err := strconv.Atoi(lines[0][len(lines[0])-1:])
	if err != nil {
		return nil, err
	}
	m1, err := strconv.Atoi(lines[1][len(lines[1])-1:])
	if err != nil {
		return nil, err
	}
	m2, err := strconv.Atoi(lines[2][len(lines[2])-1:])
	if err != nil {
		return nil, err
	}

	return func(i int) *Monkey {
		if i%d == 0 {
			return t.Monkeys[m1]
		}
		return t.Monkeys[m2]
	}, nil
}

func parseOperation(s string) (func(i int) int, error) {
	_, after, _ := strings.Cut(s, "old ")
	instructions := strings.Split(after, " ")
	op := instructions[0]
	amt := instructions[1]
	numAmt, _ := strconv.Atoi(amt)
	var isNum bool
	if amt != "old" {
		isNum = true
	}

	switch op {
	case "*":
		if isNum {
			return func(i int) int {
				return i * numAmt
			}, nil
		}
		return func(i int) int {
			return i * i
		}, nil
	case "+":
		if isNum {
			return func(i int) int {
				return i + numAmt
			}, nil
		}
		return func(i int) int {
			return i + i
		}, nil
	}

	return nil, fmt.Errorf("unable to parse %v", s)
}

func parseItems(s string) (items []*Item, err error) {
	_, numString, _ := strings.Cut(s, ": ")
	nums := strings.Split(numString, ", ")
	for _, num := range nums {
		intNum, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		items = append(items, &Item{intNum})
	}

	return items, err
}

func parseMonkeyNum(s string) (int, error) {
	snum := string(s[len(s)-2])
	num, err := strconv.Atoi(snum)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func (t *Troop) parseMonkey(lines []string) error {
	var mi int
	var m Monkey
	for i, line := range lines {
		switch i % 7 {
		// Monkey Number
		case 0:
			num, err := parseMonkeyNum(line)
			if err != nil {
				return err
			}
			mi = num
		// Starting items
		case 1:
			items, err := parseItems(line)
			if err != nil {
				return err
			}
			m.Items = items
		// Operation
		case 2:
			f, err := parseOperation(line)
			if err != nil {
				return err
			}
			m.Operation = f
		// Test
		case 3:
			f, err := t.parseTest(lines[3:])
			if err != nil {
				return err
			}
			m.Test = f
		}
	}
	t.Monkeys[mi] = &m

	return nil
}

func CreateTroop(lines []string) (*Troop, error) {
	var t Troop
	for i := range t.Monkeys {
		t.Monkeys[i] = &Monkey{}
	}

	for i := 0; i < len(lines); i += 7 {
		err := t.parseMonkey(lines[i : i+6])
		if err != nil {
			return nil, err
		}
	}

	return &t, nil
}

func Part1() (any, error) {
	lines, err := inputreader.ReadLines("inputs/day11/1sample.txt")
	if err != nil {
		return nil, err
	}

	t, err := CreateTroop(lines)
	if err != nil {
		return nil, err
	}

	t.PlayRound()
	fmt.Println(t.Monkeys[0].Items, t.Monkeys[1].Items)

	return t, nil
}

func Part2() (any, error) {
	return nil, nil
}
