package day6

import (
	"aoc2022/pkg/inputreader"
)

type window struct {
	set  map[rune]int
	dups int
}

func (w *window) Add(r rune) {
	w.set[r] += 1
	if w.set[r] > 1 {
		w.dups += 1
	}
}

func (w *window) Remove(r rune) {
	w.set[r] -= 1
	if w.set[r] > 0 {
		w.dups -= 1
	}
}

func (w *window) HasDup() bool {
	return w.dups != 0
}

func findMarker(line string) int {
	var markerIndex int

	w := &window{set: make(map[rune]int, 26)}
	for i, r := range line {
		w.Add(r)
		if i > 3 {
			w.Remove(rune(line[i-4]))
			if !w.HasDup() {
				markerIndex = i + 1
				break
			}
		}
	}

	return markerIndex
}

func Part1() (int, error) {
	lines, err := inputreader.ReadLines("inputs/day6/1.txt")
	if err != nil {
		return 0, err
	}

	markerIndex := findMarker(lines[0])
	return markerIndex, nil
}

func Part2() (int, error) {

	return 0, nil
}
