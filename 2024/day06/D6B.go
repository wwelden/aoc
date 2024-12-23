package day06

import (
	"aoc/2024/util"
	"fmt"
	"maps"
	"strings"
)

type moment struct {
	position  util.Vector
	direction util.Vector
}

func nextPosition(lines []string, position util.Vector, direction util.Vector) (util.Vector, util.Vector) {
	next := position.Add(direction)

	if isInside(lines, next) && rune(lines[next.Y][next.X]) == '#' {
		return position, direction.RotateOrigin90().Opposite()
	}

	return next, direction
}

func nextPositionWithObstacle(lines []string, position util.Vector, direction util.Vector, obstacle util.Vector) (util.Vector, util.Vector) {
	next := position.Add(direction)

	if isInside(lines, next) && (next == obstacle || rune(lines[next.Y][next.X]) == '#') {
		return position, direction.RotateOrigin90().Opposite()
	}

	return next, direction
}

func isLoop(lines []string, path util.Set[moment], position util.Vector, direction util.Vector, obstacle util.Vector) bool {
	seen := maps.Clone(path)

	for isInside(lines, position) {
		current := moment{position, direction}
		if seen.Contains(current) {
			return true
		}

		seen.Add(current)
		position, direction = nextPositionWithObstacle(lines, position, direction, obstacle)
	}

	return false
}

func walkUntilLeaves2(lines []string, position util.Vector, direction util.Vector) util.Set[util.Vector] {
	path := util.SetOf[moment]()
	stepped_on := util.SetOf[util.Vector]()
	loopObstacles := util.SetOf[util.Vector]()

	for isInside(lines, position) {
		path.Add(moment{position, direction})
		stepped_on.Add(position)

		in_front := position.Add(direction)
		if !stepped_on.Contains(in_front) && isLoop(lines, path, position, direction.RotateOrigin90().Opposite(), in_front) {
			loopObstacles.Add(in_front)
		}

		position, direction = nextPosition(lines, position, direction)
	}

	return loopObstacles
}
func testDay6PartB() {
	lines := strings.Split(testInput, "\n")
	position, direction := findStartingPostion(lines)
	obstacles := walkUntilLeaves2(lines, position, direction)

	fmt.Println("D6B Test: ", len(obstacles))
}

func SolveDay6PartB() {
	testDay6PartB()

	file := util.ReadFile("/Users/williamwelden/Developer/aoc/2024/day06/Input.txt")
	lines := strings.Split(file, "\n")

	position, direction := findStartingPostion(lines)
	obstacles := walkUntilLeaves2(lines, position, direction)

	fmt.Println("D6B: ", len(obstacles))
}
