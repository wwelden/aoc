package main

import (
	"fmt"
)

func splitLists(input [][2]int) ([]int, []int) {
	listA := sortByDex(input, 0)
	listB := sortByDex(input, 1)
	return listA, listB
}

func countOccurence(input []int, target int) int {
	count := 0
	for _, val := range input {
		if val == target {
			count++
		}
	}
	return count
}
func getSimScore(listA []int, listB []int) int {
	score := 0
	for _, val := range listA {
		score += countOccurence(listB, val) * val
	}
	return score
}

func ProblemB() {
	parsedInput := parseInput(inputFile)
	listA, listB := splitLists(parsedInput)
	count := getSimScore(listA, listB)

	fmt.Println("Ans2: ", count)
}

func main() {
	ProblemA()
	ProblemB()
}
