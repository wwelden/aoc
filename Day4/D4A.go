package main

import (
	"fmt"
	"os"
	"strings"
)

const testInput = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

var outputFile = parseInput(readInput())
var testFile = parseInput(testInput)

func readInput() string {
	inputFile, err := os.ReadFile("Input.txt")
	if err != nil {
		panic(err)
	}
	return string(inputFile)
}

func parseInput(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := make([][]rune, len(lines))

	for i, line := range lines {
		result[i] = []rune(line)
	}
	return result
}

func horizontalXMAS(input [][]rune) int {
	count := 0

	for _, line := range input {
		lineStr := string(line)
		count += strings.Count(lineStr, "XMAS")
		count += strings.Count(lineStr, "SAMX")
	}
	return count
}

func verticalXMAS(input [][]rune) int {
	count := 0

	for j := 0; j < len(input[0]); j++ {
		var vertical strings.Builder
		for i := 0; i < len(input); i++ {
			vertical.WriteRune(input[i][j])
		}
		vertStr := vertical.String()
		count += strings.Count(vertStr, "XMAS")
		count += strings.Count(vertStr, "SAMX")
	}
	return count
}

func diagonalXMAS(input [][]rune) int {
	count := 0

	for k := 0; k < len(input)*2-1; k++ {
		var diagonal strings.Builder
		for j := 0; j <= k; j++ {
			i := k - j
			if i < len(input) && j < len(input[0]) {
				diagonal.WriteRune(input[i][j])
			}
		}
		diagStr := diagonal.String()
		count += strings.Count(diagStr, "XMAS")
		count += strings.Count(diagStr, "SAMX")
	}
	return count
}

func upsideDownDiagonalXMAS(input [][]rune) int {
	count := 0

	for k := 0; k < len(input)*2-1; k++ {
		var diagonal strings.Builder
		for j := 0; j <= k; j++ {
			i := k - j
			if i < len(input) && j < len(input[0]) {
				diagonal.WriteRune(input[i][len(input[0])-j-1])
			}
		}
		diagStr := diagonal.String()
		count += strings.Count(diagStr, "XMAS")
		count += strings.Count(diagStr, "SAMX")
	}
	return count
}

var results int

func testResults() {
	results = horizontalXMAS(testFile) + verticalXMAS(testFile) + diagonalXMAS(testFile) + upsideDownDiagonalXMAS(testFile)
	fmt.Println("Test answer A:", results)
	results = ExXMAS(testFile)
	fmt.Println("Test answer B:", results)
}

func partA() {
	results = horizontalXMAS(outputFile) + verticalXMAS(outputFile) + diagonalXMAS(outputFile) + upsideDownDiagonalXMAS(outputFile)
	fmt.Println("Part A answer:", results)
}
