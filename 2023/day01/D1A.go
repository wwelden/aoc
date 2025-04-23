package day01

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const example = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
const expectedAnswer = 77

func ReadFile(path string) (string, error) {
	inputFile, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", path, err)
	}
	return string(inputFile), nil
}

func getNumMap(input string) (string, bool) {
	numberMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	input = strings.ToLower(input)
	num, ok := numberMap[input]
	return num, ok
}

func isNum(input string) (string, string) {
	length := len(input)
	if length == 0 {
		return "", ""
	}

	// Check for single digit first
	if unicode.IsDigit(rune(input[0])) {
		return string(input[0]), input[1:]
	}

	// Check for number words
	for i := 3; i <= 5; i++ {
		if length >= i {
			substring := input[:i]
			if num, ok := getNumMap(substring); ok {
				newNum, newString := isNum(input[i:])
				return num + newNum, newString
			}
		}
	}

	//No Number word or digit.
	_, newString := isNum(input[1:])
	return "", newString
}

func stringToNums(input string) []int {
	nums := make([]int, 0)
	for _, char := range input {
		num, err := strconv.Atoi(string(char))
		if err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}
func stringToNum(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0 // Handle error: return 0 or some other sentinel value
	}
	return num
}

func parse(input string) string {
	convertedString, _ := isNum(input)
	return convertedString
}

func mergeNums(input string) int {
	length := len(input)
	if length == 0 {
		return 0
	}
	var first string
	var last string
	nums := stringToNums(input)
	if len(nums) == 0 {
		return 0
	}
	first = strconv.Itoa(nums[0])
	last = strconv.Itoa(nums[len(nums)-1])

	newNum := first + last
	return stringToNum(newNum)
}

func sum(input []int) int {
	sum := 0
	for _, num := range input {
		sum += num
	}
	return sum
}

func SolveDay1PartA() {
	// file, err := ReadFile("/Users/williamwelden/Developer/aoc/2023/day01/input.txt")
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }

	lines := strings.Split(example, "\n")
	sum := 0
	for _, line := range lines {
		sum += mergeNums(parse(line))
	}
	fmt.Println("Sum:", sum)
}
