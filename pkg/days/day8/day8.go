package day8

import (
	"aoc2022/pkg/inputreader"
	"strconv"
)

type Forest struct {
	Rows    [][]*Tree
	Columns [][]*Tree
	Visible int
}

type Tree struct {
	Height  int
	Visible bool
}

func setVisible(arr []*Tree) error {
	max := -1
	for _, t := range arr {
		if t.Height > max {
			t.Visible = true
			max = t.Height
		}
	}

	max = -1
	for i := len(arr) - 1; i >= 0; i-- {
		t := arr[i]
		if t.Height > max {
			t.Visible = true
			max = t.Height
		}
	}

	return nil
}

func (f *Forest) CheckVisibility() error {
	for _, row := range f.Rows {
		setVisible(row)
	}

	for _, col := range f.Columns {
		setVisible(col)
	}

	return nil
}

func (f *Forest) SetVisible() {
	for _, row := range f.Rows {
		for _, t := range row {
			if t.Visible {
				f.Visible += 1
			}
		}
	}
}

func buildForest(lines []string) (*Forest, error) {
	var forest Forest

	var rows [][]*Tree
	var cols [][]*Tree
	for _, line := range lines {
		var row []*Tree
		for i, r := range line {
			height, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}

			t := &Tree{Height: height}
			row = append(row, t)

			if i >= len(cols) {
				cols = append(cols, []*Tree{})
			}
			cols[i] = append(cols[i], t)
		}

		rows = append(rows, row)
	}

	forest.Rows = rows
	forest.Columns = cols

	return &forest, nil
}

func Part1() (int, error) {
	lines, err := inputreader.ReadLines("inputs/day8/1.txt")
	if err != nil {
		return 0, err
	}

	forest, err := buildForest(lines)
	if err != nil {
		return 0, err
	}

	err = forest.CheckVisibility()
	if err != nil {
		return 0, err
	}

	forest.SetVisible()

	return forest.Visible, nil
}

func Part2() (int, error) {
	return 0, nil
}
