package main

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

func ProblemB() {
	results := 0
	for _, val := range parsedInput {
		if removeVal(val) {
			results++
		}
	}

	fmt.Println("Ans2: ", results)
}

func main() {
	ProblemA()
	ProblemB()
}
