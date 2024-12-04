package main

import (
	"fmt"
	"os"
	"strings"
)

//bug, got 10 extra

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
	// fmt.Println(result)
	return result

}

func horizontalXMAS(input [][]rune) int {
	count := 0

	for _, line := range input {
		// Convert line to string for easier pattern matching
		lineStr := string(line)
		// Check forward XMAS
		count += strings.Count(lineStr, "XMAS")
		// Check backward SAMX
		count += strings.Count(lineStr, "SAMX")
	}
	return count
}

func verticalXMAS(input [][]rune) int {
	count := 0

	// Create vertical strings
	for j := 0; j < len(input[0]); j++ {
		var vertical strings.Builder
		for i := 0; i < len(input); i++ {
			vertical.WriteRune(input[i][j])
		}
		vertStr := vertical.String()
		// Check forward XMAS
		count += strings.Count(vertStr, "XMAS")
		// Check backward SAMX
		count += strings.Count(vertStr, "SAMX")
	}
	return count
}

func diagonalXMAS(input [][]rune) int {
	count := 0

	// Top-left to bottom-right diagonals
	for k := 0; k < len(input)*2-1; k++ {
		var diagonal strings.Builder
		for j := 0; j <= k; j++ {
			i := k - j
			if i < len(input) && j < len(input[0]) {
				diagonal.WriteRune(input[i][j])
			}
		}
		diagStr := diagonal.String()
		// Check forward XMAS
		count += strings.Count(diagStr, "XMAS")
		// Check backward SAMX
		count += strings.Count(diagStr, "SAMX")
	}
	return count
}

func upsideDownDiagonalXMAS(input [][]rune) int {
	count := 0

	// Top-right to bottom-left diagonals
	for k := 0; k < len(input)*2-1; k++ {
		var diagonal strings.Builder
		for j := 0; j <= k; j++ {
			i := k - j
			if i < len(input) && j < len(input[0]) {
				diagonal.WriteRune(input[i][len(input[0])-j-1])
			}
		}
		diagStr := diagonal.String()
		// Check forward XMAS
		count += strings.Count(diagStr, "XMAS")
		// Check backward SAMX
		count += strings.Count(diagStr, "SAMX")
	}
	return count
}

var results int

func testResults() {
	// fmt.Println("horizontal:", horizontalXMAS(testFile))
	// fmt.Println("vertical:", verticalXMAS(testFile))
	// fmt.Println("diagonal:", diagonalXMAS(testFile))
	// fmt.Println("upside down diagonal:", upsideDownDiagonalXMAS(testFile))
	results = horizontalXMAS(testFile) + verticalXMAS(testFile) + diagonalXMAS(testFile) + upsideDownDiagonalXMAS(testFile)
	fmt.Println("Test answer A:", results)
	results = PlusXMAS(testFile) + ExXMAS(testFile)
	fmt.Println("Test answer B:", results)
}

func partA() {
	// fmt.Println("horizontal:", horizontalXMAS(outputFile))
	// fmt.Println("vertical:", verticalXMAS(outputFile))
	// fmt.Println("diagonal:", diagonalXMAS(outputFile))
	// fmt.Println("upside down diagonal:", upsideDownDiagonalXMAS(outputFile))
	results = horizontalXMAS(outputFile) + verticalXMAS(outputFile) + diagonalXMAS(outputFile) + upsideDownDiagonalXMAS(outputFile)
	fmt.Println("Part A answer:", results)
}
