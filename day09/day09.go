package main

import (
	"advent-of-code-2025/utils"
	"sort"
	"strconv"
	"strings"

	"github.com/Olegas/goaocd"
)

type Coord struct {
	y int
	x int
}

func (c Coord) String() string {
	return "(" + strconv.Itoa(c.y) + "," + strconv.Itoa(c.x) + ")"
}

func parse(input string) []Coord {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	var result []Coord
	for _, line := range lines {
		number := strings.Split(line, ",")
		x, _ := strconv.Atoi(number[0])
		y, _ := strconv.Atoi(number[1])
		c := Coord{y, x}
		result = append(result, c)
	}
	return result
}

func orientation(a, b, c Coord) int {
	var v = a.x*(b.y-c.y) + b.x*(c.y-a.y) + c.x*(a.y-b.y)
	if v < 0 {
		return -1
	}
	if v > 0 {
		return +1
	}
	return 0
}

func distanceSquared(a, b Coord) int {
	return (a.y-b.y)*(a.y-b.y) + (a.x-b.x)*(a.x-b.x)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(inputData string) int {
	var input = parse(inputData)
	var p0 = input[0]
	for _, coord := range input {
		if coord.y > p0.y || (coord.y == p0.y && coord.x < p0.x) {
			p0 = coord
		}
	}

	sort.Slice(input, func(i, j int) bool {
		var o = orientation(p0, input[i], input[j])
		if o == 0 {
			return distanceSquared(p0, input[i]) < distanceSquared(p0, input[j])
		}
		return o < 0
	})

	var stack []Coord
	stack = append(stack, p0)
	stack = append(stack, input[1])

	for i := 2; i < len(input); i++ {
		var current = input[i]
		for len(stack) >= 2 && orientation(stack[len(stack)-2], stack[len(stack)-1], current) >= 0 {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, current)
	}

	maxArea := 0
	for i, first := range stack {
		for j := i + 1; j < len(stack); j++ {
			second := stack[j]
			area := (abs(first.x-second.x) + 1) * (abs(first.y-second.y) + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func part2(inputData string) int {
	//var input = parse(inputData)
	return 0
}

const day = 9
const year = 2025

func main() {
	var testInput = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
	var input = goaocd.Input(year, day)
	utils.Assert(50, part1(testInput))
	goaocd.Submit(1, part1(input), year, day)
	utils.Assert(-1, part2(testInput))
	goaocd.Submit(2, part2(input), year, day)
}
