package main

import (
	"advent-of-code-2025/utils"
	"math"
	"strings"

	"github.com/Olegas/goaocd"
)

func parse(input string) [][]int {
	var lines = strings.Split(input, "\n")
	var banks = make([][]int, len(lines))
	for i, line := range lines {
		for _, b := range line {
			var battery = int(b - '0')
			banks[i] = append(banks[i], battery)
		}
	}
	return banks
}

func part1(inputData string) int {
	var input = parse(inputData)
	var sum = 0
	for _, bank := range input {
		var maxNum = 0
		for i, left := range bank {
			for j := i + 1; j < len(bank); j++ {
				var right = bank[j]
				var value = left*10 + right
				if value > maxNum {
					maxNum = value
				}
			}
		}
		sum += maxNum
	}
	return sum
}

func maxJoltage(used int, bank []int, value int, startIndex int, currentMax int) int {
	if used == 12 || startIndex >= len(bank) || currentMax > value*int(math.Pow(10, float64(12-used))) {
		return value
	}
	var maxNum = 0
	for i := startIndex; i < len(bank); i++ {
		maxNum = max(maxNum, maxJoltage(used+1, bank, value*10+bank[i], i+1, maxNum))
	}
	return maxNum
}

func part2(inputData string) int {
	var input = parse(inputData)
	var sum = 0
	for _, bank := range input {
		sum += maxJoltage(0, bank, 0, 0, 0)
	}
	return sum
}

const batteriesTurnedOn = 12

func maxJoltageFaster(bank []int) int {
	var number []int
	for i, b := range bank {
		if len(number)+len(bank)-i == batteriesTurnedOn || len(number) == 0 {
			number = append(number, b)
			continue
		}
		if b <= number[len(number)-1] && len(number) == batteriesTurnedOn {
			continue
		}
		for len(number) > 0 && b > number[len(number)-1] && len(number)+len(bank)-i > batteriesTurnedOn {
			number = number[:len(number)-1]
		}
		number = append(number, b)
	}
	var maxNum = 0
	for _, n := range number {
		maxNum = maxNum*10 + n
	}
	return maxNum
}

func part2Faster(inputData string) int {
	var input = parse(inputData)
	var sum = 0
	for _, bank := range input {
		sum += maxJoltageFaster(bank)
	}
	return sum
}

const day = 3
const year = 2025

func main() {
	var testInput = `987654321111111
811111111111119
234234234234278
818181911112111`

	//utils.Assert(357, part1(testInput))

	var input = goaocd.Input(year, day)
	goaocd.Submit(1, part1(input), year, day)

	utils.Assert(3121910778619, part2Faster(testInput))
	goaocd.Submit(2, part2Faster(input), year, day)

	utils.Assert(3121910778619, part2(testInput))
	goaocd.Submit(2, part2(input), year, day)
}
