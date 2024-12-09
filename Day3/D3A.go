package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const testInput = "xmul(2,4)&mul[3,7]!^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

var outputFile = readInput()

func readInput() string {
	inputFile, err := os.ReadFile("Input.txt")
	if err != nil {
		panic(err)
	}
	return string(inputFile)
}

func findValidMulOperations(input string) []int {

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	results := []int{}
	for _, match := range matches {
		x, errX := strconv.Atoi(match[1])
		y, errY := strconv.Atoi(match[2])
		if errX != nil || errY != nil {
			continue
		}
		results = append(results, x*y)
	}

	return results
}

func sumResults(results []int) int {
	sum := 0
	for _, result := range results {
		sum += result
	}
	return sum
}

func testResults() {
	results := findValidMulOperations(testInput)
	fmt.Println("Test answer A:", sumResults(results))
	results = findValidMulOperations2(testInput)
	fmt.Println("Test answer B:", sumResults(results))
}

func partA() {
	results := findValidMulOperations(outputFile)
	fmt.Println("Part A answer:", sumResults(results))
}
