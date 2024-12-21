package day05

import (
	"aoc/2024/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
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

func IsCorrectlyOrdered(pages []int, rules map[int]util.Set[int]) bool {
	seen := util.SetOf[int]()

	for _, page := range pages {
		mustBeBefore := rules[page]
		if len(mustBeBefore.Intersection(seen)) > 0 {
			return false
		}
		seen.Add(page)
	}
	return true
}

func calculateMiddleOfIncorrectlyOrderedUpdate(rules map[int]util.Set[int], update []int) int {
	if IsCorrectlyOrdered(update, rules) {
		return 0
	}

	sort.Slice(update, func(i, j int) bool {
		left := update[i]
		right := update[j]
		rule, found := rules[left]

		if found && rule.Contains(right) {
			return true
		}

		return false
	})
	return getMiddle(update)
}

func getOrderingRules(file string) map[int]util.Set[int] {
	rules := make(map[int]util.Set[int])

	for _, rule := range strings.Split(strings.TrimSpace(file), "\n") {
		rule = strings.TrimSpace(rule)
		if rule == "" {
			continue
		}

		r := strings.Split(rule, "|")
		before, err := strconv.Atoi(strings.TrimSpace(r[0]))
		if err != nil {
			panic(err)
		}

		after, err := strconv.Atoi(strings.TrimSpace(r[1]))
		if err != nil {
			panic(err)
		}

		_, found := rules[before]
		if !found {
			rules[before] = util.SetOf[int]()
		}

		rules[before].Add(after)
	}

	return rules
}

func getPageUpdates(file string) [][]int {
	lines := strings.Split(strings.TrimSpace(file), "\n")
	updates := make([][]int, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		page_strings := strings.Split(line, ",")
		pages := make([]int, 0, len(page_strings))

		for _, s := range page_strings {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}

			page, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}

			pages = append(pages, page)
		}

		updates = append(updates, pages)
	}

	return updates
}

func ParseInput(file string) (map[int]util.Set[int], [][]int) {
	parts := strings.Split(file, "\n\n")

	rules := getOrderingRules(parts[0])
	updates := getPageUpdates(parts[1])

	return rules, updates
}
func SolveDay5PartB() {
	// fixedRows := fixRows(fileInvalidRows, fileRules)

	fmt.Println(fixRows(testInvalidRows, testRules))
	results := getMiddleAll(fixRows(testInvalidRows, testRules))
	fmt.Println("D5B Test:", results)

	rules, updates := ParseInput(rawFile)

	sum := 0
	for _, update := range updates {
		sum += calculateMiddleOfIncorrectlyOrderedUpdate(rules, update)
	}

	results = sum

	fmt.Println("D5B: ", results)
}
