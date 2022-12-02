package day2

import (
	"aoc2022/pkg/inputreader"
	"strings"
)

type Move int

const (
	_ Move = iota
	Rock
	Paper
	Scissors
)

type Outcome int

const (
	Loss Outcome = 0
	Draw Outcome = 3
	Win  Outcome = 6
)

var enemyMoveMap = map[string]Move{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var myMoveMap = map[string]Move{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var RockOutcomes = map[Move]Outcome{
	Rock:     Draw,
	Paper:    Loss,
	Scissors: Win,
}

var PaperOutcomes = map[Move]Outcome{
	Rock:     Win,
	Paper:    Draw,
	Scissors: Loss,
}

var ScissorsOutcomes = map[Move]Outcome{
	Rock:     Loss,
	Paper:    Win,
	Scissors: Draw,
}

func getScore(myMove Move, enemyMove Move) int {
	var outcome Outcome

	switch myMove {
	case Rock:
		outcome = RockOutcomes[enemyMove]
	case Paper:
		outcome = PaperOutcomes[enemyMove]
	case Scissors:
		outcome = ScissorsOutcomes[enemyMove]
	}

	return int(myMove) + int(outcome)
}

func Part1() (int, error) {
	lines, err := inputreader.ReadLines("inputs/day2/1.txt")
	if err != nil {
		return 0, err
	}

	score := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		enemyMove := enemyMoveMap[fields[0]]
		myMove := myMoveMap[fields[1]]

		score += getScore(myMove, enemyMove)
	}

	return score, nil
}
