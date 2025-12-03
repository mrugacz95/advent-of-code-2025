package main

import (
	"advent-of-code-2025/utils"
	"strconv"
	"strings"

	"github.com/Olegas/goaocd"
)

func parse(input string) [][]int {
	var newLinesRemoved = strings.Join(strings.Split(input, "\n"), "")
	var ranges = strings.Split(newLinesRemoved, ",")
	var parsed = make([][]int, len(ranges))
	for i, line := range ranges {
		parsed[i] = make([]int, 2)
		bounds := strings.Split(line, "-")
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])
		parsed[i][0] = start
		parsed[i][1] = end
	}
	return parsed
}

func isValid(n int) bool {
	var s = strconv.Itoa(n)
	if len(s)%2 != 0 {
		return false
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)/2+i] {
			return false
		}
	}
	return true
}

func part1(inputData string) int {
	var input = parse(inputData)
	var validSum = 0
	for i := 0; i < len(input); i++ {
		var numRange = input[i]
		var start = numRange[0]
		var end = numRange[1]
		for j := start; j <= end; j++ {
			if isValid(j) {
				validSum += j
			}
		}
	}
	return validSum
}

func isValidPart2(n int) bool {
	var s = strconv.Itoa(n)
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i != 0 {
			continue
		}
		var matched = true
		var repeated = s[0:i]
		for j := i; j < len(s); j += i {
			if repeated != s[j:j+i] {
				matched = false
			}
		}
		if matched {
			return true
		}
	}
	return false
}

func part2(inputData string) int {
	var input = parse(inputData)
	var validSum = 0
	for i := 0; i < len(input); i++ {
		var numRange = input[i]
		var start = numRange[0]
		var end = numRange[1]
		for j := start; j <= end; j++ {
			if isValidPart2(j) {
				validSum += j
			}
		}
	}
	return validSum
}

const day = 2
const year = 2025

func main() {
	var testInput = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124`
	utils.Assert(1227775554, part1(testInput))

	var input = goaocd.Input(year, day)
	goaocd.Submit(1, part1(input), year, day)

	utils.AssertBool(true, isValidPart2(446446))
	utils.AssertBool(true, isValidPart2(38593859))
	utils.AssertBool(false, isValidPart2(824824823))

	utils.Assert(4174379265, part2(testInput))
	println(part2(input))
	goaocd.Submit(2, part2(input), year, day)
}
