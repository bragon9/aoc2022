package day7

import (
	"aoc2022/pkg/inputreader"
	"fmt"
	"strings"
)

type File struct {
	Name string
	Size int
}

type Directory struct {
	Path           []string
	Files          []File
	Subdirectories []*Directory
	Size           int
}

func handleCommand(s string) error {
	words := strings.Split(s, " ")
	fmt.Println(words)

	return nil
}

func Part1() (int, error) {
	lines, err := inputreader.ReadLines("inputs/day7/1.txt")
	if err != nil {
		return 0, err
	}

	for _, line := range lines {
		handleCommand(line)
	}

	return 0, nil
}

func Part2() (int, error) {

	return 0, nil
}
