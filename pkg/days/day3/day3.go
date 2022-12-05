package day3

import (
	"aoc2022/pkg/inputreader"
	"fmt"
)

func splitLine(line string) (string, string) {
	return line[:len(line)/2], line[len(line)/2:]
}

func makeSet(s string) map[rune]any {
	set := make(map[rune]any)
	for _, r := range s {
		set[r] = struct{}{}
	}

	return set
}

func makeScoreMap() map[rune]int {
	scoreMap := make(map[rune]int)

	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	score := 0
	for _, r := range lowerCase {
		score += 1
		scoreMap[r] = score
	}
	for _, r := range upperCase {
		score += 1
		scoreMap[r] = score
	}

	return scoreMap
}

func Part1() (int, error) {
	lines, err := inputreader.ReadLines("inputs/day3/1.txt")
	if err != nil {
		return 0, err
	}

	scoreMap := makeScoreMap()
	total := 0

	for _, line := range lines {
		c1, c2 := splitLine(line)
		c1Set := makeSet(c1)

		for _, r := range c2 {
			if _, found := c1Set[r]; found {
				total += scoreMap[r]
				break
			}
		}

	}

	return total, nil
}

func scoreGroup(arr []string) (int, error) {
	scoreMap := makeScoreMap()
	intersection := makeSet(arr[0])
	for _, line := range arr[1:3] {
		lineSet := makeSet(line)
		for r := range intersection {
			if _, found := lineSet[r]; !found {
				delete(intersection, r)
			}
		}
	}

	if len(intersection) > 1 {
		return 0, fmt.Errorf("multiple items shared \n%v", intersection)
	}

	var ans rune
	for r := range intersection {
		ans = r
	}

	return scoreMap[ans], nil
}

func Part2() (int, error) {
	lines, err := inputreader.ReadLines("inputs/day3/2.txt")
	if err != nil {
		return 0, err
	}

	total := 0
	for ptr := 0; ptr < len(lines)-2; {
		score, err := scoreGroup(lines[ptr : ptr+3])
		if err != nil {
			return 0, err
		}

		total += score
		ptr += 3
	}

	return total, nil
}
