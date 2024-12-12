package day03

import (
	"fmt"
	"regexp"
	"strconv"
)

var proceed = true

func findValidMulOperations2(input string) []int {

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	results := []int{}
	for _, match := range matches {
		if match[0] == "do()" {
			proceed = true
			continue
		} else if match[0] == "don't()" {
			proceed = false
			continue
		} else if proceed {
			x, errX := strconv.Atoi(match[1])
			y, errY := strconv.Atoi(match[2])
			if errX != nil || errY != nil {
				continue
			}
			results = append(results, x*y)
		}
	}

	return results
}

func SolveDay3PartB() {
	results := findValidMulOperations2(testInput)
	fmt.Println("D3B Test:", sumResults(results))
	results = findValidMulOperations2(outputFile)
	fmt.Println("D3B: ", sumResults(results))
}

// func main() {
// func main() {
// 	testResults()
// 	partA()
// 	partB()
// }
