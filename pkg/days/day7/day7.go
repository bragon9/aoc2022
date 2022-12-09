package day7

import (
	"aoc2022/pkg/inputreader"
	"fmt"
)

func Part1() (int, error) {
	lines, err := inputreader.ReadLines("inputs/day7/1.txt")
	if err != nil {
		return 0, err
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return 0, nil
}

func Part2() (int, error) {

	return 0, nil
}
