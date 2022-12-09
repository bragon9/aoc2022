package day5

import (
	"aoc2022/pkg/inputreader"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type move struct {
	Amount int
	From   int
	To     int
}

type craneType int

const (
	ninek craneType = iota
	ninek1
)

func getStartingStacks(lines []string) [9][]byte {
	var stacks [9][]byte
	ptr := len(lines) - 1
	for ptr >= 0 {
		line := lines[ptr]
		for i, letter := range line {
			if unicode.IsLetter(letter) {
				stacks[i/4] = append(stacks[i/4], byte(letter))
			}
		}
		ptr -= 1
	}

	return stacks
}

func getMoves(lines []string) ([]move, error) {
	var moves []move
	for _, line := range lines {
		words := strings.Fields(line)
		amount, err := strconv.Atoi(words[1])
		if err != nil {
			return []move{}, err
		}
		from, err := strconv.Atoi(words[3])
		if err != nil {
			return []move{}, err
		}
		to, err := strconv.Atoi(words[5])
		if err != nil {
			return []move{}, err
		}

		// input is 1 indexed, so subtract 1
		moves = append(moves, move{Amount: amount, From: from - 1, To: to - 1})
	}
	return moves, nil
}

func makeMove(stacks *[9][]byte, move move, ct craneType) error {
	fromStack := stacks[move.From]
	toStack := stacks[move.To]
	amt := move.Amount

	sectionToMove := fromStack[len(fromStack)-amt:]
	fromStack = fromStack[:len(fromStack)-amt]

	fmt.Println(ct)
	switch ct {
	case ninek:
		for i := 0; i < amt; i++ {
			toStack = append(toStack, sectionToMove[len(sectionToMove)-i-1])
		}
	case ninek1:
		toStack = append(toStack, sectionToMove...)
	default:
		return fmt.Errorf("unrecognized cranetype %v", ct)
	}

	stacks[move.From] = fromStack
	stacks[move.To] = toStack

	return nil
}

func getAnswer(stacks [9][]byte) string {
	var ans strings.Builder
	for _, stack := range stacks {
		ans.WriteByte(stack[len(stack)-1])
	}

	return ans.String()
}

func Part1() (string, error) {
	lines, err := inputreader.ReadLines("inputs/day5/1.txt")
	if err != nil {
		return "", err
	}

	inputSep := 0
	for i, line := range lines {
		{
			if line == "" {
				inputSep = i
				break
			}
		}
	}

	stacks := getStartingStacks(lines[:inputSep])
	moves, err := getMoves(lines[inputSep+1:])
	if err != nil {
		return "", err
	}

	for _, currMove := range moves {
		makeMove(&stacks, currMove, ninek)
		if err != nil {
			return "", err
		}
	}

	ans := getAnswer(stacks)

	return ans, nil
}

func Part2() (string, error) {
	lines, err := inputreader.ReadLines("inputs/day5/1.txt")
	if err != nil {
		return "", err
	}

	inputSep := 0
	for i, line := range lines {
		{
			if line == "" {
				inputSep = i
				break
			}
		}
	}

	stacks := getStartingStacks(lines[:inputSep])
	moves, err := getMoves(lines[inputSep+1:])
	if err != nil {
		return "", err
	}

	for _, currMove := range moves {
		err := makeMove(&stacks, currMove, ninek)
		if err != nil {
			return "", err
		}
	}

	ans := getAnswer(stacks)

	return ans, nil
}
