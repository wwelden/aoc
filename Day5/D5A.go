package main

import (
	"fmt"
	"os"
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

// var outputFile = parseInput(readInput())
var testFile = parseToTuple(testInput)
var inputFile = parseToTuple(readInput())

type tuple[T any] struct {
	a [2]T
}

func mkTuple[T any](a, b T) tuple[T] {
	return tuple[T]{[2]T{a, b}}
}
func fst[T any](t tuple[T]) T {
	return t.a[0]
}
func snd[T any](t tuple[T]) T {
	return t.a[1]
}

func readInput() string {
	inputFile, err := os.ReadFile("Input.txt")
	if err != nil {
		panic(err)
	}
	return string(inputFile)
}

func parseToTuple(input string) tuple[string] {
	lines := strings.Split(strings.TrimSpace(input), "\n\n")
	return mkTuple(lines[0], lines[1])
}

func parseRules(fileName tuple[string]) []tuple[int] {
	lines := strings.Split(fst(fileName), "\n")
	rules := make([]tuple[int], len(lines))
	for i, line := range lines {
		rule := strings.Split(line, "|")
		num1, _ := strconv.Atoi(strings.TrimSpace(rule[0]))
		num2, _ := strconv.Atoi(strings.TrimSpace(rule[1]))
		rules[i] = mkTuple(num1, num2)
	}
	return rules
}
func parsePageNums(fileName tuple[string]) [][]int {
	lines := strings.Split(snd(fileName), "\n")
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

func doesRowHaveRule(row []int, rule tuple[int]) bool {
	first := false
	second := false
	for _, num := range row {
		if num == fst(rule) {
			first = true
		}
		if num == snd(rule) {
			second = true
		}

	}
	return first && second
}

func doesRowFollowRules(row []int, rules []tuple[int]) bool {
	for _, rule := range rules {
		if doesRowHaveRule(row, rule) {
			firstIndex := -1
			secondIndex := -1
			for i, num := range row {
				if num == fst(rule) {
					firstIndex = i
				}
				if num == snd(rule) {
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
	middle := len(row) / 2
	return row[middle]
}

func getValidRows(pageNums [][]int, rules []tuple[int]) int {
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
	fmt.Println("Test answer A:", results)

	fmt.Println("Test answer B:", results)
}

func partA() {
	results = getValidRows(filePageNums, fileRules)
	fmt.Println("Part A answer:", results)
}
