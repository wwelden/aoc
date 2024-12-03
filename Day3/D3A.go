package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const testInput = ``

var inputFile = readInput()
var parsedInput = parseInput(inputFile)

func readInput() string {
	inputFile, err := os.ReadFile("Input.txt")
	if err != nil {
		panic(err)
	}
	return string(inputFile)
}

func parseInput(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := make([][]int, len(lines))

	for i, line := range lines {
		numbers := strings.Fields(line)
		result[i] = make([]int, len(numbers))
		for j, numStr := range numbers {
			num, _ := strconv.Atoi(numStr)
			result[i][j] = num
		}
	}

	return result
}

func partA() {
	fmt.Println("Part A answer:")
}
