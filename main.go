package main

import (
	"aoc2022/pkg/day1"
	"fmt"
	"log"
)

func main() {
	ans, err := day1.Part2()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(ans)
}
