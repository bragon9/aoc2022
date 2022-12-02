package main

import (
	"aoc2022/pkg/day1"
	"fmt"
	"log"
)

type Answer struct {
	Part1 any
	Part2 any
}

func getDay1() (Answer, error) {
	answer := Answer{}
	ans1, err := day1.Part1()
	if err != nil {
		return Answer{}, err
	}
	answer.Part1 = ans1
	ans2, err := day1.Part2()
	if err != nil {
		return Answer{}, err
	}
	answer.Part2 = ans2

	return answer, nil
}

func main() {
	Day1, err := getDay1()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day1: \n%+v", Day1)
}
