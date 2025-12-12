package main

import (
	"advent-of-code-2025/utils"
	"strconv"
	"strings"

	"github.com/Olegas/goaocd"
)

type Shape struct {
	content []string
	area    int
}

type Region struct {
	size       [2]int
	usedShapes []int
	area       int
}
type Day12Input struct {
	shapes  []Shape
	regions []Region
}

func parse(input string) Day12Input {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	shapes := make([]Shape, len(lines)-1)
	for i := 0; i < len(lines)-1; i++ {
		content := strings.Split(strings.TrimRight(lines[i], "\n"), "\n")
		area := 0
		for _, line := range content[1:] {
			for _, c := range line {
				if c == '#' {
					area++
				}
			}
		}
		shapes[i] = Shape{content[1:], area}
	}
	regionDefinitions := strings.Split(lines[len(lines)-1], "\n")
	regions := make([]Region, len(regionDefinitions))
	for i, definition := range regionDefinitions {
		sizeAndShapes := strings.Split(definition, ": ")
		sizeValues := strings.Split(strings.TrimSpace(sizeAndShapes[0]), "x")
		width, _ := strconv.Atoi(sizeValues[0])
		height, _ := strconv.Atoi(sizeValues[1])
		size := [2]int{width, height}
		usedShapesValues := strings.Split(sizeAndShapes[1], " ")
		usedShapes := make([]int, len(usedShapesValues))
		for j, val := range usedShapesValues {
			shapeIndex, _ := strconv.Atoi(val)
			usedShapes[j] = shapeIndex
		}
		regions[i] = Region{
			size:       size,
			usedShapes: usedShapes,
			area:       width * height,
		}
	}
	return Day12Input{shapes: shapes, regions: regions}
}

func part1(inputData string) int {
	var input = parse(inputData)
	print(len(input.shapes))
	regionsCount := 0
	for i, region := range input.regions {
		board := make([][]bool, region.size[1])
		for y := 0; y < region.size[1]; y++ {
			board[y] = make([]bool, region.size[0])
		}

		shapesArea := 0
		for shapeIndex, used := range region.usedShapes {
			shape := input.shapes[shapeIndex]
			shapesArea += shape.area * used
		}

		if shapesArea > region.area {
			println("Region", i, "cannot fit shapes: area", shapesArea, "exceeds region area", region.area)
		} else {
			println("Region", i, "can potentially fit shapes: area", shapesArea, "within region area", region.area)
			regionsCount++
		}
	}
	return regionsCount
}

const day = 12
const year = 2025

func main() {
	var testInput = `0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2`
	var input = goaocd.Input(year, day)
	utils.Assert(3, part1(testInput)) // estimation is 3
	goaocd.Submit(1, part1(input), year, day)
	utils.Assert(-1, part2(testInput))
	goaocd.Submit(2, part2(input), year, day)
}
