package day05

import (
	"aoc/2024/util"
	"fmt"
)

func getInvalidRows(pageNums [][]int, rules []util.Tuple[int]) [][]int {
	invalidRows := [][]int{}
	for _, row := range pageNums {
		if !doesRowFollowRules(row, rules) {
			invalidRows = append(invalidRows, row)
		}
	}
	return invalidRows
}

var testInvalidRows = getInvalidRows(testPageNums, testRules)
var fileInvalidRows = getInvalidRows(filePageNums, fileRules)

func getRulesForRow(row []int, rules []util.Tuple[int]) []util.Tuple[int] {
	returnRules := []util.Tuple[int]{}
	first := false
	second := false
	for _, rule := range rules {
		for _, num := range row {
			if num == util.Fst(rule) {
				first = true
			}
			if num == util.Snd(rule) {
				second = true
			}
		}
		if first && second {
			returnRules = append(returnRules, rule)
		}
		first = false
		second = false
	}
	return returnRules
}

func getRulesForNum(num int, rules []util.Tuple[int]) []util.Tuple[int] {
	returnRules := []util.Tuple[int]{}
	for _, rule := range rules {
		if num == util.Fst(rule) {
			returnRules = append(returnRules, rule)
		}
	}
	return returnRules
}

func applyRule(row []int, rule util.Tuple[int]) []int {
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
	if firstIndex > secondIndex {
		row[firstIndex], row[secondIndex] = row[secondIndex], row[firstIndex]
	}
	return row
}

func applyRules(row []int, rules []util.Tuple[int]) []int {
	rulesForRow := getRulesForRow(row, rules)
	for _, rule := range rulesForRow {
		row = applyRule(row, rule)
	}
	return row
}

func fixRows(invalidRows [][]int, rules []util.Tuple[int]) [][]int {
	// invalidRows := fileInvalidRows
	fixedRows := [][]int{}
	for _, row := range invalidRows {
		rules4Row := getRulesForRow(row, rules)
		fixedRows = append(fixedRows, applyRules(row, rules4Row))
	}
	return fixedRows
}

func getMiddleAll(rows [][]int) int {
	middle := 0
	for _, row := range rows {
		mid := getMiddle(row)
		// fmt.Println(mid)
		middle += mid
	}
	return middle
}

func SolveDay5PartB() {
	// fixedRows := fixRows(fileInvalidRows, fileRules)

	fmt.Println(fixRows(testInvalidRows, testRules))
	results := getMiddleAll(fixRows(testInvalidRows, testRules))
	fmt.Println("D5B Test:", results)
	results = getMiddleAll(fixRows(fileInvalidRows, fileRules))

	fmt.Println("D5B: ", results)
}
