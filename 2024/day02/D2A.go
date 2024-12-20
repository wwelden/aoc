package day02

import (
	"aoc/2024/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const testInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

var inputFile = util.ReadFile("/Users/williamwelden/Developer/aoc/2024/day02/Input.txt")
var parsedInput = parseInput(inputFile)

func parseInput(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := make([][]int, len(lines))

	for i, line := range lines {
		numbers := strings.Fields(line)
		result[i] = make([]int, len(numbers))
		for j, numStr := range numbers {
			num, _ := strconv.Atoi(numStr)
			result[i][j] = num
		}
	}

	return result
}

func getGaps(input []int) bool {
	val1 := input[0]
	result := true
	for i := 1; i < len(input); i++ {
		val := input[i]
		testVal := int(math.Abs(float64(val1 - val)))
		if ((notAllIncreasing(input)) && (notAllDecreasing(input))) || testVal > 3 || val1 == val {
			result = false
		}

		val1 = val
	}

	return result
}
func notAllIncreasing(input []int) bool {
	val1 := input[0]
	for _, val := range input {
		if val1 > val {
			return true
		}
		val1 = val
	}
	return false
}
func notAllDecreasing(input []int) bool {
	val1 := input[0]
	for _, val := range input {
		if val1 < val {
			return true
		}
		val1 = val
	}
	return false
}

func ProblemA() {
	results := 0
	for _, val := range parsedInput {
		if getGaps(val) {
			results++
		}
	}

	fmt.Println("Ans1: ", results)
}

func testResultsA() {
	testInput := parseInput(testInput)
	results := 0
	for _, val := range testInput {
		if getGaps(val) {
			results++
		}
	}
	fmt.Println("D2A Test:", results)
}

func SolveDay2PartA() {
	testResultsA()
	results := 0
	for _, val := range parsedInput {
		if getGaps(val) {
			results++
		}
	}

	fmt.Println("D2A: ", results)
}

// func main() {
// 	ProblemA()
// }
