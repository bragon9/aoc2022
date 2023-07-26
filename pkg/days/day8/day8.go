package day8

import (
	"aoc2022/pkg/inputreader"
	"strconv"
)

type Forest struct {
	Rows           [][]*Tree
	Columns        [][]*Tree
	Visible        int
	MaxScenicScore int
}

type Tree struct {
	Height      int
	Visible     bool
	ScenicScore int
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

func (f *Forest) SetVisiblity() {
	for _, row := range f.Rows {
		for _, t := range row {
			if t.Visible {
				f.Visible += 1
			}
		}
	}
}

func getYAxisScore(arr []*Tree, i int) int {
	southScore := 1
	for y := i + 1; y < len(arr)-1; y++ {
		if arr[y].Height >= arr[i].Height {
			break
		}
		southScore += 1
	}

	northScore := 1
	for y := i - 1; y > 0; y-- {
		if arr[y].Height >= arr[i].Height {
			break
		}
		northScore += 1
	}

	return southScore * northScore
}

func getXAxisScore(arr []*Tree, i int) int {
	eastScore := 1
	for x := i + 1; x < len(arr)-1; x++ {
		if arr[x].Height >= arr[i].Height {
			break
		}
		eastScore += 1
	}

	westScore := 1
	for x := i - 1; x > 0; x-- {
		if arr[x].Height >= arr[i].Height {
			break
		}
		westScore += 1
	}

	return eastScore * westScore
}

func (f *Forest) SetScenicScores() {
	maxScore := 0
	for i := 1; i < len(f.Rows[0])-1; i++ {
		for j := 1; j < len(f.Columns[0])-1; j++ {
			t := f.Rows[i][j]
			t.ScenicScore = getXAxisScore(f.Rows[i], j) * getYAxisScore(f.Columns[j], i)
			if t.ScenicScore > maxScore {
				maxScore = t.ScenicScore
			}
		}
	}

	f.MaxScenicScore = maxScore
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

	forest.SetVisiblity()

	return forest.Visible, nil
}

func Part2() (int, error) {
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

	forest.SetVisiblity()
	forest.SetScenicScores()

	return forest.MaxScenicScore, nil
}
