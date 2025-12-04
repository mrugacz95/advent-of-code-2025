package main

import (
	"advent-of-code-2025/utils"
	"strings"

	"github.com/Olegas/goaocd"
)

func parse(input string) []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func getAccessibleRolls(grid []string) [][]int {
	var accessibe [][]int
	neighbours := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	for y, line := range grid {
		for x, c := range line {
			if c == '.' {
				continue
			}
			neighboursCount := 0
			for _, neighbour := range neighbours {
				dx := neighbour[0]
				dy := neighbour[1]
				if x+dx < 0 || x+dx >= len(line) {
					continue
				}
				if y+dy < 0 || y+dy >= len(grid) {
					continue
				}
				if grid[y+dy][x+dx] == '@' {
					neighboursCount++
				}
			}
			if neighboursCount < 4 {
				accessibe = append(accessibe, []int{y, x})
			}
		}
	}
	return accessibe
}

func part1(inputData string) int {
	input := parse(inputData)
	ans := len(getAccessibleRolls(input))
	return ans
}

func part2(inputData string) int {
	var input = parse(inputData)
	ans := 0
	for {
		accessible := getAccessibleRolls(input)
		for i := 0; i < len(accessible); i++ {
			y, x := accessible[i][0], accessible[i][1]
			input[y] = input[y][:x] + "." + input[y][x+1:]
		}
		if len(accessible) == 0 {
			break
		}
		ans += len(accessible)
	}
	return ans
}

// ..@@.@@@@.
// ..@.@@@@.

const day = 4
const year = 2025

func main() {
	var testInput = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	var input = goaocd.Input(year, day)
	utils.Assert(13, part1(testInput))
	goaocd.Submit(1, part1(input), year, day)
	utils.Assert(43, part2(testInput))
	goaocd.Submit(2, part2(input), year, day)
}

// 1395 too low
