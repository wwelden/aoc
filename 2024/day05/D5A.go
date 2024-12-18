package day05

import (
	"aoc/2024/util"
	"fmt"
	"strconv"
	"strings"
)

const testInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

var testFile = parseToTuple(testInput)
var inputFile = parseToTuple(util.ReadFile("/Users/williamwelden/Developer/aoc/2024/day05/Input.txt"))
var rawFile = util.ReadFile("/Users/williamwelden/Developer/aoc/2024/day05/Input.txt")

// type Tuple[T any] struct {
// 	a [2]T
// }

// func MkTuple[T any](a, b T) Tuple[T] {
// 	return Tuple[T]{[2]T{a, b}}
// }
// func Fst[T any](t Tuple[T]) T {
// 	return t.a[0]
// }
// func Snd[T any](t Tuple[T]) T {
// 	return t.a[1]
// }

func parseToTuple(input string) util.Tuple[string] {
	lines := strings.Split(strings.TrimSpace(input), "\n\n")
	return util.MkTuple(lines[0], lines[1])
}

func parseRules(fileName util.Tuple[string]) []util.Tuple[int] {
	lines := strings.Split(util.Fst(fileName), "\n")
	rules := make([]util.Tuple[int], len(lines))
	for i, line := range lines {
		rule := strings.Split(line, "|")
		num1, _ := strconv.Atoi(strings.TrimSpace(rule[0]))
		num2, _ := strconv.Atoi(strings.TrimSpace(rule[1]))
		rules[i] = util.MkTuple(num1, num2)
	}
	return rules
}

func parsePageNums(fileName util.Tuple[string]) [][]int {
	lines := strings.Split(util.Snd(fileName), "\n")
	nums := make([][]int, len(lines))
	for i, line := range lines {
		nums[i] = make([]int, len(strings.Split(line, ",")))
		for j, num := range strings.Split(line, ",") {
			nums[i][j], _ = strconv.Atoi(num)
		}
	}
	return nums
}

var testRules = parseRules(testFile)
var testPageNums = parsePageNums(testFile)
var fileRules = parseRules(inputFile)
var filePageNums = parsePageNums(inputFile)

func doesRowHaveRule(row []int, rule util.Tuple[int]) bool {
	first := false
	second := false
	for _, num := range row {
		if num == util.Fst(rule) {
			first = true
		}
		if num == util.Snd(rule) {
			second = true
		}

	}
	return first && second
}

func doesRowFollowRules(row []int, rules []util.Tuple[int]) bool {
	for _, rule := range rules {
		if doesRowHaveRule(row, rule) {
			firstIndex := -1
			secondIndex := -1
			for i, num := range row {
				if num == util.Fst(rule) {
					firstIndex = i
				}
				if num == util.Snd(rule) {
					secondIndex = i
				}
			}
			if firstIndex > secondIndex || firstIndex == -1 || secondIndex == -1 {
				return false
			}
		}
	}
	return true
}

func getMiddle(row []int) int {
	if len(row)%2 == 0 {
		panic("Cannot get middle of even length row")
	}
	middle := (len(row) - 1) / 2
	return row[middle]
}

func getValidRows(pageNums [][]int, rules []util.Tuple[int]) int {
	sum := 0
	for _, row := range pageNums {
		if doesRowFollowRules(row, rules) {
			sum += getMiddle(row)
		}
	}
	return sum
}

var results int

func testResults() {
	results = getValidRows(testPageNums, testRules)
	fmt.Println("D5A Test:", results)
}

func SolveDay5PartA() {
	testResults()
	results = getValidRows(filePageNums, fileRules)
	fmt.Println("D5A: ", results)
}
