package main

import (
	"advent-of-code-2025/utils"
	"regexp"
	"strconv"
	"strings"

	"github.com/Olegas/goaocd"
)

type InputDay6 struct {
	numbers   [][]int
	operators []rune
}

func parse(input string) InputDay6 {
	var re = regexp.MustCompile(" +")
	var lines = strings.Split(strings.TrimRight(input, "\n"), "\n")
	var numbers = make([][]int, len(lines)-1)
	for i := 0; i < len(lines)-1; i++ {
		line := strings.Trim(lines[i], " ")
		var columns = re.Split(line, -1)
		for _, num := range columns {
			value, _ := strconv.Atoi(num)
			numbers[i] = append(numbers[i], value)
		}
	}
	var lastLine = lines[len(lines)-1]
	var columns = re.Split(strings.Trim(lastLine, " "), -1)
	var operators = make([]rune, len(columns))
	for i := 0; i < len(columns); i++ {
		operators[i] = rune(columns[i][0])
	}
	return InputDay6{numbers, operators}
}
func calculate(input InputDay6) int {
	var ans = 0
	for i := 0; i < len(input.numbers[0]); i++ {
		var op = input.operators[i]
		var columnValue = input.numbers[0][i]
		for j := 1; j < len(input.numbers); j++ {
			if op == '+' {
				columnValue += input.numbers[j][i]
			} else {
				columnValue *= input.numbers[j][i]
			}
		}
		ans += columnValue
	}
	return ans
}

func part1(inputData string) int {
	input := parse(inputData)
	return calculate(input)
}

func part2(inputData string) int {
	parsed := parse(inputData)
	splitInput := strings.Split(strings.TrimRight(inputData, "\n"), "\n")
	x := 0
	ans := 0
	for i := 0; i < len(parsed.numbers[0]); i++ { // iterate group
		var columnResult = 0
		first := true
		var longest = parsed.numbers[0][i] // find longest to determine number width
		for j := 1; j < len(parsed.numbers); j++ {
			longest = max(longest, parsed.numbers[j][i])
		}
		longestNumLen := len(strconv.Itoa(longest))
		for col := x + longestNumLen - 1; col >= x; col-- { // iterate columns of numbers
			newNumber := ""
			for row := 0; row < len(parsed.numbers); row++ { // construct number from rows
				if col >= len(splitInput[row]) {
					newNumber += " "
					continue
				}
				newNumber += string(splitInput[row][col])
			}
			value, _ := strconv.Atoi(strings.TrimSpace(newNumber))
			if first { // calculate number depending on operator
				columnResult = value
				first = false
			} else {
				if parsed.operators[i] == '+' {
					columnResult += value
				} else {
					columnResult *= value
				}
			}
		}
		ans += columnResult
		x += longestNumLen + 1
	}
	return ans
}

const day = 6
const year = 2025

func main() {
	var testInput = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
	var input = goaocd.Input(year, day)
	utils.Assert(4277556, part1(testInput))
	goaocd.Submit(1, part1(input), year, day)
	utils.Assert(3263827, part2(testInput))
	goaocd.Submit(2, part2(input), year, day)
}
