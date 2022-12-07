package day4

import (
	"aoc2022/pkg/inputreader"
	"fmt"
	"strconv"
	"strings"
)

type window struct {
	Start int
	Stop  int
}

func convertAssignmentToWindow(assignment string) (window, error) {
	var window window

	startAndStop := strings.Split(assignment, "-")
	start, err := strconv.Atoi(startAndStop[0])
	if err != nil {
		return window, err
	}
	stop, err := strconv.Atoi(startAndStop[1])
	if err != nil {
		return window, err
	}

	window.Start = start
	window.Stop = stop

	return window, nil
}

func hasCompleteOverlap(assignments []string) (bool, error) {
	if len(assignments) != 2 {
		return false, fmt.Errorf("%v assignments passed in.  Expected 2", len(assignments))
	}

	window1, err := convertAssignmentToWindow(assignments[0])
	if err != nil {
		return false, err
	}
	window2, err := convertAssignmentToWindow(assignments[1])
	if err != nil {
		return false, err
	}

	if window1.Start >= window2.Start && window1.Stop <= window2.Stop {
		return true, nil
	}

	if window2.Start >= window1.Start && window2.Stop <= window1.Stop {
		return true, nil
	}

	return false, nil
}

func hasAnyOverlap(assignments []string) (bool, error) {
	if len(assignments) != 2 {
		return false, fmt.Errorf("%v assignments passed in.  Expected 2", len(assignments))
	}

	window1, err := convertAssignmentToWindow(assignments[0])
	if err != nil {
		return false, err
	}
	window2, err := convertAssignmentToWindow(assignments[1])
	if err != nil {
		return false, err
	}

	if window1.Start >= window2.Start && window1.Start <= window2.Stop {
		return true, nil
	}

	if window1.Stop >= window2.Start && window1.Stop <= window2.Stop {
		return true, nil
	}

	if window2.Start >= window1.Start && window2.Start <= window1.Stop {
		return true, nil
	}

	if window2.Stop >= window1.Start && window2.Stop <= window1.Stop {
		return true, nil
	}

	return false, nil
}

func Part1() (int, error) {
	lines, err := inputreader.ReadLines("inputs/day4/1.txt")
	if err != nil {
		return 0, err
	}

	total := 0
	for _, line := range lines {
		splitLines := strings.Split(line, ",")
		overlap, err := hasCompleteOverlap(splitLines)
		if err != nil {
			return 0, err
		}

		if overlap {
			total += 1
		}
	}

	return total, nil
}

func Part2() (int, error) {
	lines, err := inputreader.ReadLines("inputs/day4/1.txt")
	if err != nil {
		return 0, err
	}

	total := 0
	for _, line := range lines {
		splitLines := strings.Split(line, ",")
		overlap, err := hasAnyOverlap(splitLines)
		if err != nil {
			return 0, err
		}

		if overlap {
			total += 1
		}
	}

	return total, nil
}
