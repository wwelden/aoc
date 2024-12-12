package day03

import (
	"aoc/2024/util"
	"fmt"
	"regexp"
	"strconv"
)

const testInput = "xmul(2,4)&mul[3,7]!^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

var outputFile = util.ReadFile("/Users/williamwelden/Developer/aoc/2024/day03/Input.txt")

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

func SolveDay3PartA() {
	results := findValidMulOperations(testInput)
	fmt.Println("D3A Test:", sumResults(results))
	results = findValidMulOperations(outputFile)
	fmt.Println("D3A: ", sumResults(results))
}
