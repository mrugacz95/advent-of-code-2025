package main

import (
	"advent-of-code-2025/utils"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/Olegas/goaocd"
)

type Position struct {
	x int
	y int
	z int
}

func (a Position) String() string {
	return "(" + strconv.Itoa(a.x) + "," + strconv.Itoa(a.y) + "," + strconv.Itoa(a.z) + ")"
}

type Closest struct {
	first    int
	second   int
	distance int
}

func parse(input string) []Position {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	var positions = make([]Position, len(lines))
	for i, line := range lines {
		pos := strings.Split(line, ",")
		x, _ := strconv.Atoi(pos[0])
		y, _ := strconv.Atoi(pos[1])
		z, _ := strconv.Atoi(pos[2])
		positions[i] = Position{x: x, y: y, z: z}
	}
	return positions
}

func findParent(parent *[]int, i int) int {
	if (*parent)[i] != i {
		(*parent)[i] = findParent(parent, (*parent)[i])
	}
	return (*parent)[i]
}

func union(parent *[]int, i int, j int) {
	pi := findParent(parent, i)
	pj := findParent(parent, j)
	if pi != pj {
		(*parent)[pj] = pi
	}
}

func distance(a Position, b Position) int {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return dx*dx + dy*dy + dz*dz
}

func initParents(n int) []int {
	parents := make([]int, n)
	for i := range parents {
		parents[i] = i
	}
	return parents
}

func calcDistances(positions []Position) [][]int {
	var distances = make([][]int, len(positions))
	for i := 0; i < len(positions); i++ {
		distances[i] = make([]int, len(positions))
		for j := 0; j < len(positions); j++ {
			if i == j {
				distances[i][j] = math.MaxInt
				continue
			}
			distances[i][j] = distance(positions[i], positions[j])
		}
	}
	return distances
}

func closest(arr [][]int) Closest {
	minVal := math.MaxInt
	first := -1
	second := -1
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			}
			if arr[i][j] < minVal {
				minVal = arr[i][j]
				first = i
				second = j
			}
		}
	}
	return Closest{first: first, second: second, distance: minVal}
}

func part1(inputData string, rounds int) int {
	var positions = parse(inputData)

	var parents = initParents(len(positions))
	var distances = calcDistances(positions)

	for i := 0; i < rounds; i++ {
		current := closest(distances)

		println("Closest:", positions[current.first].String(), " ", current.first, " x ", positions[current.second].String(), " ", current.second, "Distance:", current.distance)
		union(&parents, current.first, current.second)
		println("New parents ", parents[current.first], " and ", parents[current.second])

		distances[current.first][current.second] = math.MaxInt
		distances[current.second][current.first] = math.MaxInt
	}

	groups := make(map[int]int)
	for i := 0; i < len(parents); i++ {
		root := findParent(&parents, i)
		if _, exists := groups[root]; !exists {
			groups[root] = 0
		}
		groups[root] += 1
	}
	var sizes []int
	for _, size := range groups {
		sizes = append(sizes, size)
	}
	sort.Ints(sizes)
	ans := 1
	for i := len(sizes) - 1; i >= len(sizes)-3; i-- {
		ans *= sizes[i]
	}
	return ans
}

func part2(inputData string) int {
	//var input = parse(inputData)
	return 0
}

const day = 8
const year = 2025

func main() {
	var testInput = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`
	var input = goaocd.Input(year, day)
	utils.Assert(40, part1(testInput, 10))
	goaocd.Submit(1, part1(input, 1000), year, day)
	utils.Assert(-1, part2(testInput))
	goaocd.Submit(2, part2(input), year, day)
}
