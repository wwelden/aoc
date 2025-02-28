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
			if c != '.' && c != '#' { // Assuming '.' is empty space and '#' is an antinode
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

			// For each potential antinode position
			for x := 0; x < width; x++ {
				for y := 0; y < height; y++ {
					antinode := Position{X: x, Y: y}

					// Skip if the antinode is at the same position as either antenna
					if (antinode.X == a1.X && antinode.Y == a1.Y) ||
						(antinode.X == a2.X && antinode.Y == a2.Y) {
						continue
					}

					// Check if the three points are collinear
					if !isCollinear(a1, antinode, a2) {
						continue
					}

					// Calculate distances (squared to avoid floating point)
					d1 := distanceSquared(antinode, a1)
					d2 := distanceSquared(antinode, a2)

					// Check if one distance is twice the other
					if d1 == 4*d2 || d2 == 4*d1 {
						antinodes[antinode] = true
					}
				}
			}
		}
	}

	return antinodes
}

// Check if three points are collinear
func isCollinear(p1, p2, p3 Position) bool {
	// Using the formula: (y2-y1)*(x3-x1) == (y3-y1)*(x2-x1)
	return (p2.Y-p1.Y)*(p3.X-p1.X) == (p3.Y-p1.Y)*(p2.X-p1.X)
}

// Calculate the squared distance between two points
func distanceSquared(p1, p2 Position) int {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return dx*dx + dy*dy
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
