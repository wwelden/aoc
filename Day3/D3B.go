package main

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

func partB() {
	results := findValidMulOperations2(outputFile)
	fmt.Println("Part B answer:", sumResults(results))
}

func main() {
	testResults()
	partA()
	partB()
}
