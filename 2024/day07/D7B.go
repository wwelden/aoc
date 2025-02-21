package day07

import (
	"aoc/2024/util"
	"fmt"
	"strconv"
)

type Node3 struct {
	value int
	mult  *Node3
	add   *Node3
	comb  *Node3
}

func combine(a int, b int) int {
	strA := fmt.Sprint(a)
	strB := fmt.Sprint(b)

	ret, errAtoi := strconv.Atoi((strA + strB))
	if errAtoi != nil {
		fmt.Println("NAN")
	}
	return ret
}

func buildTree2(lst []int) *Node3 {
	if len(lst) == 0 {
		return nil
	}

	// Root node starts with the first value
	root := &Node3{value: lst[0]}

	if len(lst) > 1 {
		// Compute new values
		multVal := root.value * lst[1]
		addVal := root.value + lst[1]
		combVal := combine(root.value, lst[1])

		// Assign computed values to the child nodes
		root.mult = buildTree2(append([]int{multVal}, lst[2:]...))
		root.add = buildTree2(append([]int{addVal}, lst[2:]...))
		root.comb = buildTree2(append([]int{combVal}, lst[2:]...))
	}

	return root
}
func checkNodes2(node *Node3, target int) bool {
	if node == nil {
		return false
	}
	if node.mult == nil && node.add == nil && node.comb == nil {
		return node.value == target
	}
	return checkNodes2(node.mult, target) || checkNodes2(node.add, target) || checkNodes2(node.comb, target)
}
func isWorkingVal2(input equation) bool {
	orgVal := input.testVal
	rootNode := buildTree2(input.operands)
	return checkNodes2(rootNode, orgVal)
}
func sumWorkingVals2(input string) int {
	equations := parse(input)
	sum := 0

	for _, eq := range equations {
		if isWorkingVal2(eq) {
			sum += eq.testVal
		}
	}

	return sum
}

func testDay7PartB() {
	result := sumWorkingVals2(testInput)
	fmt.Println("Test Sum of working values:", result)
}

func SolveDay7PartB() {
	testDay7PartB()
	fmt.Println("Sum of working values:", sumWorkingVals2(util.ReadFile("/Users/williamwelden/Developer/aoc/2024/day07/Input.txt")))
}
