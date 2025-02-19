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

// Node represents a node in the binary tree.
type Node struct {
	value int
	mult  *Node
	add   *Node
}

// buildTree constructs the binary tree from the given list.
func buildTree(lst []int) *Node {
	if len(lst) == 0 {
		return nil
	}

	// Root node starts with the first value
	root := &Node{value: lst[0]}

	if len(lst) > 1 {
		// Compute new values
		multVal := root.value * lst[1]
		addVal := root.value + lst[1]

		// Assign computed values to the child nodes
		root.mult = buildTree(append([]int{multVal}, lst[2:]...))
		root.add = buildTree(append([]int{addVal}, lst[2:]...))
	}

	return root
}

func printNode(node Node) {
	if node.mult == nil {
		fmt.Println("End node: ", node.value)
		fmt.Println("")
	} else {
		fmt.Println("Node: ", node.value)
		fmt.Println("Mult node: ", node.mult.value)
		fmt.Println("Add node: ", node.add.value)
		fmt.Println("")
		printNode(*node.mult)
		printNode(*node.add)
	}
}

func printTree(lst []int) {
	tree := buildTree(lst)
	printNode(*tree)
}
func printPaths(node *Node, path []int, ops []string) {
	if node == nil {
		return
	}

	// Add the current node value to the path
	path = append(path, node.value)

	// If we are at a leaf node, print the path
	if node.mult == nil && node.add == nil {
		// Build the equation string
		equation := fmt.Sprintf("%d", path[0])
		for i := 1; i < len(path); i++ {
			equation += fmt.Sprintf(" %s %d", ops[i-1], path[i])
		}
		fmt.Println(equation, "=", path[len(path)-1])
		return
	}

	// Recur for left (Multiplication)
	if node.mult != nil {
		printPaths(node.mult, append([]int(nil), path...), append(ops, "*"))
	}

	// Recur for right (Addition)
	if node.add != nil {
		printPaths(node.add, append([]int(nil), path...), append(ops, "+"))
	}
}

// Wrapper function
func printTreePaths(lst []int) {
	tree := buildTree(lst)
	printPaths(tree, []int{}, []string{})
}

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

	test := []int{5, 10, 15, 20, 25}
	fmt.Println("test: ", test)

	fmt.Println("hello: ", parsedFile[0].testVal)
	fmt.Println("mult: ", mult(parsedFile[0].operands))
	// fmt.Println("div: ", div(test))
	fmt.Println("add: ", add(parsedFile[0].operands))
	// fmt.Println("sub: ", sub(parsedFile[0].operands))

	printTreePaths(test)
}
