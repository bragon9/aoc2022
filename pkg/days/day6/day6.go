package day6

import (
	"aoc2022/pkg/inputreader"
)

const (
	packetLength  = 4
	messageLength = 14
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

func findMarker(line string, length int) int {
	var markerIndex int

	w := &window{set: make(map[rune]int, 26)}
	for i, r := range line {
		w.Add(r)
		if i > length-1 {
			w.Remove(rune(line[i-length]))
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

	markerIndex := findMarker(lines[0], packetLength)
	return markerIndex, nil
}

func Part2() (int, error) {
	lines, err := inputreader.ReadLines("inputs/day6/1.txt")
	if err != nil {
		return 0, err
	}

	markerIndex := findMarker(lines[0], messageLength)
	return markerIndex, nil
}
