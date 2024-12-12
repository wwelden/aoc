package main

import (
	"aoc/2024/day01"
	"aoc/2024/day02"
	"aoc/2024/day03"
	"aoc/2024/day04"
	"aoc/2024/day05"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		panic("Not enough arguments - supply day and problem numbers")
	}
	day := os.Args[1]
	problem := os.Args[2]

	switch day {
	case "1":
		switch problem {
		case "1":
			day01.SolveDay1PartA()
		case "2":
			day01.SolveDay1PartB()
		}
	case "2":
		switch problem {
		case "1":
			day02.SolveDay2PartA()
		case "2":
			day02.SolveDay2PartB()
		}
	case "3":
		switch problem {
		case "1":
			day03.SolveDay3PartA()
		case "2":
			day03.SolveDay3PartB()
		}
	case "4":
		switch problem {
		case "1":
			day04.SolveDay4PartA()
		case "2":
			day04.SolveDay4PartB()
		}
	case "5":
		switch problem {
		case "1":
			day05.SolveDay5PartA()
		case "2":
			day05.SolveDay5PartB()
		}
	default:
		fmt.Println("Invalid input")
	}
}
