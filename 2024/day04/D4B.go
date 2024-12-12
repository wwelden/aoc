package day04

import (
	"fmt"
)

func PlusXMAS(input [][]rune) int {
	count := 0
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			if input[i][j] == 'A' {
				if input[i-1][j] == 'M' && input[i][j+1] == 'S' &&
					input[i+1][j] == 'M' && input[i][j-1] == 'S' {
					count++
				}
				if input[i-1][j] == 'M' && input[i][j+1] == 'M' &&
					input[i+1][j] == 'S' && input[i][j-1] == 'S' {
					count++
				}
				if input[i-1][j] == 'S' && input[i][j+1] == 'S' &&
					input[i+1][j] == 'M' && input[i][j-1] == 'M' {
					count++
				}
				if input[i-1][j] == 'S' && input[i][j+1] == 'M' &&
					input[i+1][j] == 'S' && input[i][j-1] == 'M' {
					count++
				}
			}
		}
	}
	return count
}

func ExXMAS(input [][]rune) int {
	count := 0
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			if input[i][j] == 'A' {
				if input[i-1][j+1] == 'M' && input[i-1][j-1] == 'S' &&
					input[i+1][j+1] == 'M' && input[i+1][j-1] == 'S' {
					count++
				}
				if input[i-1][j+1] == 'M' && input[i-1][j-1] == 'M' &&
					input[i+1][j+1] == 'S' && input[i+1][j-1] == 'S' {
					count++
				}
				if input[i-1][j+1] == 'S' && input[i-1][j-1] == 'S' &&
					input[i+1][j+1] == 'M' && input[i+1][j-1] == 'M' {
					count++
				}
				if input[i-1][j+1] == 'S' && input[i-1][j-1] == 'M' &&
					input[i+1][j+1] == 'S' && input[i+1][j-1] == 'M' {
					count++
				}
			}
		}
	}
	return count
}

func SolveDay4PartB() {
	results = ExXMAS(testFile)
	fmt.Println("D4B Test:", results)
	results = ExXMAS(outputFile)
	fmt.Println("D4B: ", results)
}
