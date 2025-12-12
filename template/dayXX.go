package main

import (
	"advent-of-code-2025/utils"
	"strings"

	"github.com/Olegas/goaocd"
)

func parse(input string) []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func part1(inputData string) int {
	//var input = parse(inputData)

	return 0
}

func part2(inputData string) int {
	//var input = parse(inputData)
	return 0
}

const day = -1
const year = 2025

func main() {
	var testInput = `123`
	var input = goaocd.Input(year, day)
	utils.Assert(-1, part1(testInput))
	goaocd.Submit(1, part1(input), year, day)
	utils.Assert(-1, part2(testInput))
	goaocd.Submit(2, part2(input), year, day)
}
