package day06

import (
	"aoc/2024/util"
	"fmt"
	"strings"
)

func findStartingPostion(lines []string) (util.Vector, util.Vector) {
	for y, line := range lines {
		for x, c := range line {
			if c == '^' {
				return util.Vector{X: x, Y: y}, util.Vector{X: 0, Y: -1}
			}
		}
	}

	panic("Could not find \"^\" (starting position) in map.")
}

func isInside(lines []string, position util.Vector) bool {
	if position.Y < 0 || position.Y > len(lines)-1 {
		return false
	}

	if position.X < 0 || position.X > len(lines[position.Y])-1 {
		return false
	}

	return true
}

func NextMove(lines []string, position util.Vector, direction util.Vector, directions_checked int) (util.Vector, util.Vector) {
	next := position.Add(direction)

	if isInside(lines, next) && rune(lines[next.Y][next.X]) == '#' {
		if directions_checked == 3 {
			panic(fmt.Sprintln("I'm trapped!", position))
		}

		return NextMove(lines, position, direction.RotateOrigin90().Opposite(), directions_checked+1)
	}

	return next, direction
}

func walkUntilLeaves(lines []string, position util.Vector, direction util.Vector) util.Set[util.Vector] {
	uniqueLocations := util.SetOf[util.Vector]()

	for isInside(lines, position) {
		uniqueLocations.Add(position)

		position, direction = NextMove(lines, position, direction, 0)
	}

	return uniqueLocations
}

const testInput = `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func testDay6PartA() {
	lines := strings.Split(testInput, "\n")
	position, direction := findStartingPostion(lines)
	uniqueLocations := walkUntilLeaves(lines, position, direction)

	fmt.Println("D6A Test: ", len(uniqueLocations))
}

func SolveDay6PartA() {
	testDay6PartA()
	file := util.ReadFile("/Users/williamwelden/Developer/aoc/2024/day06/Input.txt")
	lines := strings.Split(file, "\n")

	position, direction := findStartingPostion(lines)
	uniqueLocations := walkUntilLeaves(lines, position, direction)

	fmt.Println("D6A: ", len(uniqueLocations))
}
