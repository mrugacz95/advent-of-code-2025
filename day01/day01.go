package main

import (
	"strconv"

	"github.com/Olegas/goaocd"
)

// Test solution for 2021 Day 1 Part 1
func main() {
	var input = goaocd.Lines(2021, 1)
	var lastNum, _ = strconv.Atoi(input[0])
	var sum = 0
	for i := 1; i < len(input); i++ {
		var n, _ = strconv.Atoi(input[i])
		if n > lastNum {
			sum += 1
		}
		lastNum = n
	}
	goaocd.Submit(1, sum, 2021, 1)
}
