package main

import (
	"advent-of-code-2025/utils"
	"strings"
	"time"

	"github.com/Olegas/goaocd"
)

func parse(input string) []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

type Pair struct {
	x int
	y int
}

func part1(inputData string) int {
	var input = parse(inputData)
	var start = strings.IndexAny(input[0], "S")
	var queue []Pair
	visited := make(map[Pair]bool)
	queue = append(queue, Pair{y: 0, x: start})
	splitCount := 0
	for len(queue) > 0 {
		var current = queue[0]
		queue = queue[1:]

		for current.y+1 < len(input) && current.x >= 0 && current.x < len(input[0]) {
			current = Pair{y: current.y + 1, x: current.x}
			if visited[current] {
				break
			}
			if input[current.y][current.x] == '^' {
				visited[current] = true
				splitCount++
				queue = append(queue, Pair{y: current.y, x: current.x - 1})
				queue = append(queue, Pair{y: current.y, x: current.x + 1})
				break
			}
		}
	}
	return splitCount
}

func part2Slow(inputData string) int {
	var input = parse(inputData)
	var start = strings.IndexAny(input[0], "S")
	var queue []Pair
	queue = append(queue, Pair{y: 0, x: start})
	timelinesCount := 0
	for len(queue) > 0 {
		var current = queue[0]
		queue = queue[1:]

		for current.x >= 0 && current.x < len(input[0]) {
			current = Pair{y: current.y + 1, x: current.x}
			if current.y == len(input) {
				timelinesCount++
				break
			}
			if input[current.y][current.x] == '^' {
				queue = append(queue, Pair{y: current.y, x: current.x - 1})
				queue = append(queue, Pair{y: current.y, x: current.x + 1})
				break
			}
		}
	}
	return timelinesCount

}

func nextSplit(board []string, start Pair) Pair {
	for start.y+1 < len(board) && start.x >= 0 && start.x < len(board[0]) {
		start = Pair{y: start.y + 1, x: start.x}
		if board[start.y][start.x] == '^' {
			return Pair{y: start.y, x: start.x}
		}
	}
	return start
}

func countTimelines(board []string, start Pair, memo map[Pair]int) int {
	if val, ok := memo[start]; ok {
		return val
	}
	if start.y >= len(board)-1 {
		return 1
	}
	left := nextSplit(board, Pair{y: start.y, x: start.x - 1})
	right := nextSplit(board, Pair{y: start.y, x: start.x + 1})
	timelines := countTimelines(board, left, memo) + countTimelines(board, right, memo)
	memo[start] = timelines
	return timelines
}

func part2Fast(inputData string) int {
	var input = parse(inputData)
	var start = strings.IndexAny(input[0], "S")
	memo := make(map[Pair]int)
	return countTimelines(input, Pair{y: 0, x: start}, memo)
}

const day = 7
const year = 2025

func main() {
	testInput := `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`
	var input = goaocd.Input(year, day)
	utils.Assert(21, part1(testInput))
	goaocd.Submit(1, part1(input), year, day)

	t := time.Now()
	utils.Assert(40, part2Slow(testInput))
	println("Part 2 slow solution:", time.Since(t))

	smallTestInput := `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............`

	utils.Assert(8, part2Slow(smallTestInput))
	utils.Assert(8, part2Fast(smallTestInput))

	t = time.Now()
	utils.Assert(40, part2Fast(testInput))
	println("Part 2 fast solution:", time.Since(t))

	goaocd.Submit(2, part2Fast(input), year, day)
}

// too high 1685
