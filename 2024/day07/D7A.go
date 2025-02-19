package day07

import (
	"aoc/2024/util"
	"fmt"
	"strconv"
	"strings"
)

const testInput = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

const Operations = `*/+-`

func mult(lst []int) int {
	if len(lst) > 1 {
		return lst[0] * mult(lst[1:])
	} else {
		return lst[0]
	}
}
func div(lst []int) int {
	if len(lst) > 1 {
		return lst[0] / div(lst[1:])
	} else {
		return lst[0]
	}
}
func add(lst []int) int {
	if len(lst) > 1 {
		return lst[0] + add(lst[1:])
	} else {
		return lst[0]
	}
}
func sub(lst []int) int {
	if len(lst) > 1 {
		return lst[0] - sub(lst[1:])
	} else {
		return lst[0]
	}
}

type equation struct {
	testVal  int
	operands []int
}

type node struct {
	mult  *node
	add   *node
	value int
}

// func (eq *equation) getTestVal(input int) {
// 	return eq.testVal
// }

func parse(input string) []equation {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var equations []equation

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue // Skip malformed lines
		}

		testVal, err := strconv.Atoi(parts[0])
		if err != nil {
			continue // Skip invalid test values
		}

		operandStrs := strings.Fields(parts[1]) // Split remaining numbers
		var operands []int
		for _, numStr := range operandStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				continue // Skip invalid numbers
			}
			operands = append(operands, num)
		}

		equations = append(equations, equation{
			testVal:  testVal,
			operands: operands,
		})
	}

	return equations
}

func testDay7PartA() {

}

func SolveDay7PartA() {
	testDay7PartA()
	file := util.ReadFile("/Users/williamwelden/Developer/aoc/2024/day07/Input.txt")
	parsedFile := parse(file)

	test := make([]int, 5)
	for i := range test {
		test[i] = (i + 1) * 5
	}
	fmt.Println("test: ", test)

	fmt.Println("hello: ", parsedFile[0].testVal)
	fmt.Println("mult: ", mult(parsedFile[0].operands))
	// fmt.Println("div: ", div(test))
	fmt.Println("add: ", add(parsedFile[0].operands))
	// fmt.Println("sub: ", sub(parsedFile[0].operands))
}
