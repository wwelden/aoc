package day08

import (
	"aoc/2024/util"
	"strings"
)

const testInput = `......#....#
...#....0...
....#0....#.
..#....0....
....0....#..
.#....A.....
...#........
#......#....
........A...
.........A..
..........#.
..........#.`

type Grid [][]rune

func parse(input string) Grid {
	lines := strings.Split(input, "\n")
	grid := make(Grid, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func testDay8PartA() {

}

func SolveDay8PartA() {
	testDay8PartA()
	file := util.ReadFile("/Users/williamwelden/Developer/aoc/2024/day08/Input.txt")
	grid := parse(file)

}
