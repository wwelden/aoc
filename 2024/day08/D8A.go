package day08

import (
	"aoc/2024/util"
	"fmt"
	"strings"
)

const testInput = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

type Grid [][]rune
type Position struct {
	X, Y int
}

func parse(input string) Grid {
	lines := strings.Split(input, "\n")
	grid := make(Grid, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

// Find all positions of antennas with a specific frequency
func findAntennas(grid Grid) map[rune][]Position {
	antennas := make(map[rune][]Position)
	for y, row := range grid {
		for x, c := range row {
			if c != '.' && c != '#' { // Skip empty space and antinode markers
				antennas[c] = append(antennas[c], Position{x, y})
			}
		}
	}
	return antennas
}

// Find all antinodes for a given set of antennas with the same frequency
func findAntinodes(antennas []Position, width, height int) map[Position]bool {
	antinodes := make(map[Position]bool)

	// For each pair of antennas
	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			a1 := antennas[i]
			a2 := antennas[j]

			// Calculate the two antinode positions directly
			// First antinode: 2*a2 - a1 (where point is twice as far from a1 as from a2)
			antinode1X := 2*a2.X - a1.X
			antinode1Y := 2*a2.Y - a1.Y

			// Second antinode: 2*a1 - a2 (where point is twice as far from a2 as from a1)
			antinode2X := 2*a1.X - a2.X
			antinode2Y := 2*a1.Y - a2.Y

			// Check if the first antinode is within bounds
			if antinode1X >= 0 && antinode1X < width && antinode1Y >= 0 && antinode1Y < height {
				antinodes[Position{antinode1X, antinode1Y}] = true
			}

			// Check if the second antinode is within bounds
			if antinode2X >= 0 && antinode2X < width && antinode2Y >= 0 && antinode2Y < height {
				antinodes[Position{antinode2X, antinode2Y}] = true
			}
		}
	}

	return antinodes
}

func solve(grid Grid) int {
	height := len(grid)
	width := len(grid[0])

	// Find all antennas grouped by frequency
	antennasByFreq := findAntennas(grid)

	// Track all antinodes
	allAntinodes := make(map[Position]bool)

	// For each frequency, find all antinodes
	for _, antennas := range antennasByFreq {
		if len(antennas) >= 2 { // Need at least 2 antennas to create antinodes
			antinodes := findAntinodes(antennas, width, height)
			for pos := range antinodes {
				allAntinodes[pos] = true
			}
		}
	}

	return len(allAntinodes)
}

func testDay8PartA() {
	grid := parse(testInput)
	result := solve(grid)
	expected := 14 // From the problem statement

	if result == expected {
		fmt.Println("Test passed!")
	} else {
		fmt.Printf("Test failed! Got %d, expected %d\n", result, expected)
	}
}

func SolveDay8PartA() {
	testDay8PartA()
	file := util.ReadFile("/Users/williamwelden/Developer/aoc/2024/day08/Input.txt")
	grid := parse(file)

	answer := solve(grid)
	fmt.Printf("Day 8 Part A Answer: %d\n", answer)
}
