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

func parse(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	// var results [][]string
	// for line := range lines {
	// }
	return lines
}

func testDay8PartA() {

}

func SolveDay8PartA() {
	testDay8PartA()
	file := util.ReadFile("/Users/williamwelden/Developer/aoc/2024/day08/Input.txt")
	data := parse(file)

}
