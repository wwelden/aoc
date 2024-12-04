package main

import (
	"fmt"
)

func PlusXMAS(input [][]rune) int {
	count := 0

	// Check plus pattern (+)
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			// Center must be 'A'
			if input[i][j] == 'A' {
				// Check all four possible plus patterns:
				// M in center-up, S in arms
				if input[i-1][j] == 'M' && input[i][j+1] == 'M' && input[i+1][j] == 'S' && input[i][j-1] == 'S' {
					count++
				}
				// M in center-right, S in arms
				if input[i-1][j] == 'S' && input[i][j+1] == 'S' && input[i+1][j] == 'M' && input[i][j-1] == 'M' {
					count++
				}
				// M in center-down, S in arms
				if input[i-1][j] == 'S' && input[i][j+1] == 'M' && input[i+1][j] == 'S' && input[i][j-1] == 'M' {
					count++
				}
				// M in center-left, S in arms
				if input[i-1][j] == 'M' && input[i][j+1] == 'S' && input[i+1][j] == 'M' && input[i][j-1] == 'M' {
					count++
				}
			}
		}
	}
	return count
}

func ExXMAS(input [][]rune) int {
	count := 0

	// Check X pattern
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			// Center must be 'A'
			if input[i][j] == 'A' {
				// Check all four possible X patterns:
				// M in top-right, S in other diagonals
				if input[i-1][j+1] == 'M' && input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S' && input[i+1][j-1] == 'S' {
					count++
				}
				// M in bottom-right, S in other diagonals
				if input[i-1][j+1] == 'S' && input[i-1][j-1] == 'S' && input[i+1][j+1] == 'M' && input[i+1][j-1] == 'M' {
					count++
				}
				// M in bottom-left, S in other diagonals
				if input[i-1][j+1] == 'S' && input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S' && input[i+1][j-1] == 'M' {
					count++
				}
				// M in top-left, S in other diagonals
				if input[i-1][j+1] == 'M' && input[i-1][j-1] == 'S' && input[i+1][j+1] == 'M' && input[i+1][j-1] == 'S' {
					count++
				}
			}
		}
	}
	return count
}

func partB() {
	fmt.Println("Plus XMAS:", PlusXMAS(outputFile))
	fmt.Println("Ex XMAS:", ExXMAS(outputFile))
	results := PlusXMAS(outputFile) + ExXMAS(outputFile)
	fmt.Println("Part B answer:", results)
}

func main() {
	testResults()
	partA()
	partB()
}
