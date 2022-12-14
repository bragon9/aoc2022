package main

import (
	"aoc2022/pkg/days/day1"
	"aoc2022/pkg/days/day2"
	"aoc2022/pkg/days/day3"
	"aoc2022/pkg/days/day4"
	"aoc2022/pkg/days/day5"
	"aoc2022/pkg/days/day6"
	"aoc2022/pkg/days/day7"
	"fmt"
	"log"
)

type Answer struct {
	Part1 any
	Part2 any
}

func getDay1() (Answer, error) {
	var answer Answer
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

func getDay2() (Answer, error) {
	var answer Answer
	ans1, err := day2.Part1()
	if err != nil {
		return Answer{}, err
	}
	answer.Part1 = ans1
	ans2, err := day2.Part2()
	if err != nil {
		return Answer{}, err
	}
	answer.Part2 = ans2

	return answer, nil
}

func getDay3() (Answer, error) {
	var answer Answer
	ans1, err := day3.Part1()
	if err != nil {
		return Answer{}, err
	}
	answer.Part1 = ans1
	ans2, err := day3.Part2()
	if err != nil {
		return Answer{}, err
	}
	answer.Part2 = ans2

	return answer, nil
}

func getDay4() (Answer, error) {
	var answer Answer
	ans1, err := day4.Part1()
	if err != nil {
		return Answer{}, err
	}
	answer.Part1 = ans1
	ans2, err := day4.Part2()
	if err != nil {
		return Answer{}, err
	}
	answer.Part2 = ans2

	return answer, nil
}

func getDay5() (Answer, error) {
	var answer Answer
	ans1, err := day5.Part1()
	if err != nil {
		return Answer{}, err
	}
	answer.Part1 = ans1
	ans2, err := day5.Part2()
	if err != nil {
		return Answer{}, err
	}
	answer.Part2 = ans2

	return answer, nil
}

func getDay6() (Answer, error) {
	var answer Answer
	ans1, err := day6.Part1()
	if err != nil {
		return Answer{}, err
	}
	answer.Part1 = ans1
	ans2, err := day6.Part2()
	if err != nil {
		return Answer{}, err
	}
	answer.Part2 = ans2

	return answer, nil
}

func getDay7() (Answer, error) {
	var answer Answer
	ans1, err := day7.Part1()
	if err != nil {
		return Answer{}, err
	}
	answer.Part1 = ans1
	ans2, err := day7.Part2()
	if err != nil {
		return Answer{}, err
	}
	answer.Part2 = ans2

	return answer, nil
}

func main() {
	// Day1, err := getDay1()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Day1: \n%+v", Day1)

	// Day2, err := getDay2()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Day2: \n%+v", Day2)

	// Day3, err := getDay3()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Day3: \n%+v", Day3)

	// Day4, err := getDay4()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Day4: \n%+v", Day4)

	// Day5, err := getDay5()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Day5: \n%+v", Day5)

	// Day6, err := getDay6()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Day6: \n%+v", Day6)

	Day7, err := getDay7()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day7: \n%+v", Day7)
}
