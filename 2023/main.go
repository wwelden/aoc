package main

import (
	"aoc/2023/day01"
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
		}
	}
}
