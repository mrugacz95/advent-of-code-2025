package main

import (
	"advent-of-code-2025/utils"
	"sort"
	"strconv"
	"strings"

	"github.com/Olegas/goaocd"
)

type Interval struct {
	start int
	end   int
}
type Ingredients struct {
	intervals []Interval
	ids       []int
}

func parse(input string) Ingredients {
	ingredients := strings.Split(input, "\n\n")
	var intervals []Interval
	for _, ingredient := range strings.Split(ingredients[0], "\n") {
		lines := strings.Split(ingredient, "-")
		start, _ := strconv.Atoi(lines[0])
		end, _ := strconv.Atoi(lines[1])
		intervals = append(intervals, Interval{start, end})
	}
	var ids []int
	for _, value := range strings.Split(ingredients[1], "\n") {
		id, _ := strconv.Atoi(value)
		ids = append(ids, id)
	}
	return Ingredients{intervals: intervals, ids: ids}
}

func part1(inputData string) int {
	var input = parse(inputData)
	ans := 0
	for _, id := range input.ids {
		fresh := false
		for _, interval := range input.intervals {
			if id >= interval.start && id <= interval.end {
				fresh = true
				break
			}
		}
		if fresh {
			ans += 1
		}
	}
	return ans
}

func (interval Interval) overlaps(other Interval) bool {
	return interval.start <= other.end && other.start <= interval.end
}

func (interval Interval) join(other Interval) Interval {
	newStart := min(interval.start, other.start)
	newEnd := max(interval.end, other.end)
	return Interval{newStart, newEnd}
}

func mergeIntervals(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})
	last := intervals[0]
	var merged []Interval
	for i := 1; i < len(intervals); i++ {
		var current = intervals[i]
		if last.overlaps(current) {
			last = last.join(current)
		} else {
			merged = append(merged, last)
			last = current
		}
	}
	merged = append(merged, last)
	return merged
}

func part2(inputData string) int {
	var input = parse(inputData)
	ans := 0
	merged := mergeIntervals(input.intervals)
	for _, interval := range merged {
		ans += interval.end - interval.start + 1
	}
	return ans
}

const day = 5
const year = 2025

func main() {
	var testInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	var input = goaocd.Input(year, day)
	utils.Assert(3, part1(testInput))
	goaocd.Submit(1, part1(input), year, day)
	utils.Assert(14, part2(testInput))
	goaocd.Submit(2, part2(input), year, day)
}

// too low 387
