package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/Olegas/goaocd"
)

func parse(input string) []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func part1(inputData string) int {
	var input = parse(inputData)
	var ans = 0
	var pos = 50
	for _, l := range input {
		var dir, num = l[0], l[1:]
		var step, _ = strconv.Atoi(num)
		if dir == 'R' {
			pos += step
		}
		if dir == 'L' {
			pos -= step
		}
		pos = (pos + 100) % 100
		if pos == 0 {
			ans++
		}
	}
	return ans
}

func part2(inputData string) int {
	var input = parse(inputData)
	var ans = 0
	var pos = 50
	for _, l := range input {
		var dir, num = l[0], l[1:]
		var step, _ = strconv.Atoi(num)
		var delta = 1
		if dir == 'L' {
			delta = -1
		}
		for range step {
			pos += delta
			pos = (pos + 100) % 100
			if pos == 0 {
				ans += 1
			}
		}
		fmt.Print(pos, " ", ans, "\n")
	}
	return ans
}

const day = 1
const year = 2025

func main() {
	var testInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	var input = goaocd.Input(year, day)
	utils.Assert(part1(testInput), 3)
	goaocd.Submit(1, part1(input), year, day)
	utils.Assert(part2(testInput), 6)
	goaocd.Submit(2, part1(input), year, day)
	part2(input)
}
