package day2

import (
	"aoc2022/pkg/inputreader"
	"strings"
)

type move int

const (
	_ move = iota
	Rock
	Paper
	Scissors
)

type outcome int

const (
	Loss outcome = 0
	Draw outcome = 3
	Win  outcome = 6
)

var enemyMoveMap = map[string]move{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var myMoveMap = map[string]move{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var rockOutcomes = map[move]outcome{
	Rock:     Draw,
	Paper:    Loss,
	Scissors: Win,
}

var paperOutcomes = map[move]outcome{
	Rock:     Win,
	Paper:    Draw,
	Scissors: Loss,
}

var scissorsOutcomes = map[move]outcome{
	Rock:     Loss,
	Paper:    Win,
	Scissors: Draw,
}

func getOutcome(enemyMove move, myMove move) outcome {
	var outcome outcome
	switch myMove {
	case Rock:
		outcome = rockOutcomes[enemyMove]
	case Paper:
		outcome = paperOutcomes[enemyMove]
	case Scissors:
		outcome = scissorsOutcomes[enemyMove]
	}

	return outcome
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

		outcome := getOutcome(enemyMove, myMove)

		score += int(myMove) + int(outcome)
	}

	return score, nil
}

var outcomeMap = map[string]outcome{
	"X": Loss,
	"Y": Draw,
	"Z": Win,
}

var moveToWinMap = map[move]move{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

var moveToLoseMap = map[move]move{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

func chooseMove(enemyMove move, outcome outcome) move {
	var myMove move
	switch outcome {
	case Loss:
		myMove = moveToLoseMap[enemyMove]
	case Win:
		myMove = moveToWinMap[enemyMove]
	case Draw:
		myMove = enemyMove
	}

	return myMove
}

func Part2() (int, error) {
	lines, err := inputreader.ReadLines("inputs/day2/1.txt")
	if err != nil {
		return 0, err
	}

	score := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		enemyMove := enemyMoveMap[fields[0]]
		outcome := outcomeMap[fields[1]]

		myMove := chooseMove(enemyMove, outcome)

		score += int(outcome) + int(myMove)
	}

	return score, nil
}
