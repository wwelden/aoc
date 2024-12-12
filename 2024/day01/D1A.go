package day01

import (
	"aoc/2024/util"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

const testInput = `3 4
4 3
2 5
1 3
3 9
3 3`

var inputFile = util.ReadFile("/Users/williamwelden/Developer/aoc/2024/day01/Input.txt")
var parsedInput = parseInput(inputFile)

func parseInput(input string) [][2]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := make([][2]int, len(lines))

	for i, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) == 2 {
			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])
			result[i] = [2]int{num1, num2}
		}
	}

	return result
}

func sortByDex(input [][2]int, dex int) []int {
	result := make([]int, len(input))
	for i, val := range input {
		result[i] = val[dex]
	}
	sort.Ints(result)
	return result
}

func getDifferenceList(inputA []int, inputB []int) []int {
	if len(inputA) != len(inputB) {
		return nil
	}
	results := make([]int, len(inputA))
	for i, val := range inputA {
		results[i] = int(math.Abs(float64(val - inputB[i])))
	}
	return results
}

func getSumOfList(input []int) int {
	sum := 0
	for _, val := range input {
		sum += val
	}
	return sum
}

func testResultsA() {
	testInput := parseInput(testInput)
	listA := sortByDex(testInput, 0)
	listB := sortByDex(testInput, 1)
	diffList := getDifferenceList(listA, listB)
	sum := getSumOfList(diffList)
	fmt.Println("D1A Test:", sum)
}

func SolveDay1PartA() {
	testResultsA()
	listA := sortByDex(parsedInput, 0)
	listB := sortByDex(parsedInput, 1)
	diffList := getDifferenceList(listA, listB)
	sum := getSumOfList(diffList)
	fmt.Println("D1A: ", sum)
}
