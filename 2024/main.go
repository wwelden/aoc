package main

import (
	"aoc/2024/day01"
	"aoc/2024/day02"
	"aoc/2024/day03"
	"aoc/2024/day04"
	"aoc/2024/day05"
	"aoc/2024/day06"
	"aoc/2024/day07"
	"aoc/2024/day08"
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
	case "6":
		switch problem {
		case "1":
			day06.SolveDay6PartA()
		case "2":
			day06.SolveDay6PartB()
		}
	case "7":
		switch problem {
		case "1":
			day07.SolveDay7PartA()
		case "2":
			day07.SolveDay7PartB()
		}
	case "8":
		switch problem {
		case "1":
			day08.SolveDay8PartA()
		case "2":
			day08.SolveDay8PartB()
		}
	case "9":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	case "10":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	case "11":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	case "12":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	case "13":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	case "14":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	case "15":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	case "16":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	case "17":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	case "18":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	case "19":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	case "20":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	case "21":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	case "22":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	case "23":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	case "24":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	case "25":
		switch problem {
		case "1":
			fmt.Println("problem not started yet")
		case "2":
			fmt.Println("problem not started yet")
		}
	default:
		fmt.Println("Invalid input")
	}
}
