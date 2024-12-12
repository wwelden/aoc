package day02

import (
	"fmt"
)

func removeVal(input []int) bool {
	result := false
	for i := 0; i < len(input); i++ {
		// Create temporary slice without element at index i
		temp := make([]int, 0)
		temp = append(temp, input[:i]...)
		temp = append(temp, input[i+1:]...)

		if getGaps(temp) {
			result = true
			break
		}
	}
	return result
}

func testResultsB() {
	testInput := parseInput(testInput)
	results := 0
	for _, val := range testInput {
		if removeVal(val) {
			results++
		}
	}
	fmt.Println("D2B Test:", results)
}

func SolveDay2PartB() {
	testResultsB()
	results := 0
	for _, val := range parsedInput {
		if removeVal(val) {
			results++
		}
	}

	fmt.Println("D2B: ", results)
}
