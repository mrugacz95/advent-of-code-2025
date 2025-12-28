package main

import (
	"advent-of-code-2025/utils"
	"sort"
	"strconv"
	"strings"

	"github.com/Olegas/goaocd"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type Coord struct {
	y int
	x int
}

func (c Coord) String() string {
	return "(y:" + strconv.Itoa(c.y) + ", x:" + strconv.Itoa(c.x) + ")"
}

func parse(input string) []Coord {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	var result []Coord
	for _, line := range lines {
		number := strings.Split(line, ",")
		x, _ := strconv.Atoi(number[0])
		y, _ := strconv.Atoi(number[1])
		c := Coord{y, x}
		result = append(result, c)
	}
	return result
}

func orientation(a, b, c Coord) int {
	var v = a.x*(b.y-c.y) + b.x*(c.y-a.y) + c.x*(a.y-b.y)
	if v < 0 {
		return -1
	}
	if v > 0 {
		return +1
	}
	return 0
}

func distanceSquared(a, b Coord) int {
	return (a.y-b.y)*(a.y-b.y) + (a.x-b.x)*(a.x-b.x)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findConvexHaul(input []Coord) []Coord {
	var p0 = input[0]
	for _, coord := range input {
		if coord.y > p0.y || (coord.y == p0.y && coord.x < p0.x) {
			p0 = coord
		}
	}

	sort.Slice(input, func(i, j int) bool {
		var o = orientation(p0, input[i], input[j])
		if o == 0 {
			return distanceSquared(p0, input[i]) < distanceSquared(p0, input[j])
		}
		return o < 0
	})

	var stack []Coord
	stack = append(stack, p0)
	stack = append(stack, input[1])

	for i := 2; i < len(input); i++ {
		var current = input[i]
		for len(stack) >= 2 && orientation(stack[len(stack)-2], stack[len(stack)-1], current) >= 0 {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, current)
	}
	return stack
}

func part1(inputData string) int {
	var input = parse(inputData)
	var points = findConvexHaul(input)

	maxArea := 0
	for i, first := range points {
		for j := i + 1; j < len(points); j++ {
			second := points[j]
			area := (abs(first.x-second.x) + 1) * (abs(first.y-second.y) + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func plotSimpleData(inputData string, fileName string) {
	var input = parse(inputData)
	p := plot.New()

	pts := make(plotter.XYs, len(input)+1)
	for i, coord := range input {
		pts[i].X = float64(coord.x)
		pts[i].Y = float64(coord.y)
	}

	pts[len(input)].X = float64(input[0].x)
	pts[len(input)].Y = float64(input[0].y)

	_ = plotutil.AddLinePoints(p, "Points", pts)

	_ = p.Save(10*vg.Inch, 10*vg.Inch, fileName)
}

func plotData(inputData string, fileName string, name string, lines plotter.XYs) {
	var input = parse(inputData)
	p := plot.New()

	pts := make(plotter.XYs, len(input)+1)
	for i, coord := range input {
		pts[i].X = float64(coord.x)
		pts[i].Y = float64(coord.y)
	}

	pts[len(input)].X = float64(input[0].x)
	pts[len(input)].Y = float64(input[0].y)

	var shape = make(plotter.XYs, 1)
	shape[0].Y = float64(48492)
	shape[0].X = float64(94916)

	_ = plotutil.AddLinePoints(p, "Points", pts, name, lines)
	_ = plotutil.AddScatters(p, "X", shape)

	_ = p.Save(10*vg.Inch, 10*vg.Inch, fileName)
}

func plotRect(inputData string, fileName string, first Coord, second Coord) {

	xmin := min(first.x, second.x)
	xmax := max(first.x, second.x)
	ymin := min(first.y, second.y)
	ymax := max(first.y, second.y)

	plotData(inputData, fileName, "rectangle", plotter.XYs{
		{X: float64(xmin), Y: float64(ymin)},
		{X: float64(xmin), Y: float64(ymax)},
		{X: float64(xmax), Y: float64(ymax)},
		{X: float64(xmax), Y: float64(ymin)},
		{X: float64(xmin), Y: float64(ymin)},
	})
}

type Interval struct {
	top, bottom Coord
}

func (i Interval) intersects(other Interval) bool {
	return i.bottom.y >= other.top.y && other.bottom.y >= i.top.y
}

func (i Interval) contains(coord Coord) bool {
	return i.top.y <= coord.y && coord.y <= i.bottom.y
}

func (i Interval) length() int {
	return abs(i.bottom.y-i.top.y) + 1
}

func (i Interval) String() string {
	return "[" + i.top.String() + " - " + i.bottom.String() + "]"
}

func part2(inputData string) int { // doesnt work, need to fix
	var input = parse(inputData)
	sorted := make([]Coord, len(input))
	copy(sorted, input)
	sort.Slice(sorted, func(i, j int) bool {
		var lhs, rhs = sorted[i], sorted[j]
		return lhs.y > rhs.y || (lhs.y == rhs.y && lhs.x < rhs.x)
	})

	verticalLines := collectVerticalLines(input)

	activeHeadlines := make([]Interval, 0)
	interactionPoints := make([]int, 0)
	for x := range verticalLines {
		interactionPoints = append(interactionPoints, x)
	}
	sort.Ints(interactionPoints)

	maxArea := 0
	for _, x := range interactionPoints {
		line := verticalLines[x]
		println("Processing line at x=", x, " from ", line[0].String(), " to ", line[1].String())
		top := line[0]
		bottom := line[1]

		if top.y > bottom.y {
			panic("Wrong line orientation")
		}

		for _, coord := range line {
			for _, headline := range activeHeadlines {
				if headline.contains(coord) {
					area := calcArea(headline, coord)
					println("New area found with headline ", headline.String(), " and coord ", coord.String(), " = ", area)
					maxArea = max(maxArea, area)
				}
			}
		}

		newActiveHeadlines := make([]Interval, 0)
		for _, headline := range activeHeadlines {
			ended := false
			for _, coord := range line {
				if headline.contains(coord) && coord.y != headline.top.y && coord.y != headline.bottom.y {
					ended = true
					break
				}
			}
			if !ended {
				newActiveHeadlines = append(newActiveHeadlines, headline)
			}
		}

		extended := make([]Interval, 0)
		for _, line := range newActiveHeadlines {
			if line.top.y == bottom.y {
				extended = append(extended, Interval{top: top, bottom: Coord{y: line.bottom.y, x: top.x}})
			}
			if line.bottom.y == top.y {
				extended = append(extended, Interval{top: Coord{y: line.top.y, x: bottom.x}, bottom: bottom})
			}
		}

		newActiveHeadlines = append(newActiveHeadlines, Interval{top: top, bottom: bottom})
		newActiveHeadlines = append(newActiveHeadlines, extended...)

		activeHeadlines = newActiveHeadlines

		println("  Active headlines: ")
		for _, headline := range activeHeadlines {
			print(headline.String())
			print(", ")
		}
		println()

	}
	return maxArea
}

func calcArea(headline Interval, coord Coord) int {
	var topArea = (abs(headline.top.y-coord.y) + 1) * (abs(headline.top.x-coord.x) + 1)
	var botArea = (abs(headline.bottom.y-coord.y) + 1) * (abs(headline.bottom.x-coord.x) + 1)
	return max(topArea, botArea)
}

func calcCoordArea(first Coord, second Coord) int {
	return (abs(first.y-second.y) + 1) * (abs(first.x-second.x) + 1)
}

func collectVerticalLines(input []Coord) map[int][]Coord {
	lines := make(map[int][]Coord)
	for _, coord := range input {
		if _, ok := lines[coord.x]; !ok {
			lines[coord.x] = []Coord{coord}
		} else {
			if lines[coord.x][0].y < coord.y {
				lines[coord.x] = append(lines[coord.x], coord)
			} else {
				lines[coord.x] = append([]Coord{coord}, lines[coord.x]...)
			}
		}
	}
	for x := range lines {
		if len(lines[x]) != 2 { // check input data
			panic("Line " + strconv.Itoa(x) + " has more than two points")
		}
	}
	return lines
}

func isInside(first Coord, second Coord, tested Coord) bool {
	xmin := min(first.x, second.x)
	xmax := max(first.x, second.x)
	ymin := min(first.y, second.y)
	ymax := max(first.y, second.y)
	return xmin < tested.x && tested.x < xmax &&
		ymin < tested.y && tested.y < ymax
}

func checkPossibleArea(corner Coord, candidates []Coord, points []Coord) (int, Coord) {
	maxArea := 0
	best := corner
	for _, candidate := range candidates {
		pointInside := false
		for _, point := range points {
			if isInside(corner, candidate, point) && point != corner && point != candidate {
				pointInside = true
				break
			}
		}
		if !pointInside {
			area := calcCoordArea(corner, candidate)
			println(area, candidate.x, candidate.y)
			if area > maxArea {
				maxArea = area
				best = candidate
			}
		}
	}
	return maxArea, best
}

func lazyPart2(inputData string) int {
	input := parse(inputData)

	maxLength := 0
	var points = make([]Coord, 0)
	for i := range len(input) {
		first := input[i]
		var second Coord
		if i+1 < len(input) {
			second = input[i+1]
		} else {
			second = input[0]
		}
		length := abs(first.x-second.x) + abs(first.y-second.y)
		if length > maxLength {
			maxLength = length
			points = []Coord{first, second}
		}
	}

	position := points[0].x
	verticalLines := collectVerticalLines(input)
	line := verticalLines[position]
	top := line[1]
	bottom := line[0]

	if top.y < bottom.y {
		panic("Wrong line orientation")
	}

	println("top", top.String(), " bot", bottom.String())

	plotData(inputData, "day09/found.png", "checked", plotter.XYs{
		plotter.XY{
			Y: float64(top.y), X: float64(top.x),
		},
		plotter.XY{
			Y: float64(bottom.y), X: float64(bottom.x),
		},
	})

	points = make([]Coord, 0)
	var topCandidates = make(plotter.XYs, 0)
	candidates := make([]Coord, 0)
	for _, coord := range input {
		if coord.y > top.y {
			topCandidates = append(topCandidates, plotter.XY{Y: float64(coord.y), X: float64(coord.x)})
			candidates = append(candidates, coord)
		}
	}
	if len(candidates) == 0 {
		panic("No candidates found for top")
	}
	println("candidates for top:", len(candidates))
	topCandidates = append(topCandidates, plotter.XY{Y: float64(top.y), X: float64(top.x)})
	maxArea, best := checkPossibleArea(top, candidates, candidates)

	println("top best", best.String(), " area=", maxArea)
	plotRect(inputData, "day09/toprect.png", top, best)

	plotData(inputData, "day09/top.png", "checked", topCandidates)

	var botChecked = make(plotter.XYs, 0)
	candidates = make([]Coord, 0)
	for _, coord := range input {
		if coord.y < bottom.y {
			botChecked = append(botChecked, plotter.XY{Y: float64(coord.y), X: float64(coord.x)})
			candidates = append(candidates, coord)
		}
	}

	if len(candidates) == 0 {
		panic("No candidates found for bot")
	}

	println("candidates for bot:", len(candidates))
	botChecked = append(botChecked, plotter.XY{Y: float64(bottom.y), X: float64(bottom.x)})
	plotData(inputData, "day09/bot.png", "checked", botChecked)

	botMax, best := checkPossibleArea(bottom, candidates, candidates)
	println("bot best", best.String(), " area=", botMax)
	plotRect(inputData, "day09/botrect.png", bottom, best)

	maxArea = max(maxArea, botMax)
	return maxArea
}

const day = 9
const year = 2025

func main() {
	var testInput = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
	var input = goaocd.Input(year, day)
	utils.Assert(50, part1(testInput))
	goaocd.Submit(1, part1(input), year, day)

	plotSimpleData(testInput, "day09/part1.png")
	plotSimpleData(input, "day09/part2.png")

	//utils.Assert(24, part2(testInput))

	//	secondInput := `6,0
	//6,3
	//2,3
	//2,5
	//11,5
	//11,8
	//8,8
	//8,11
	//12,11
	//12,13
	//17,13
	//17,10
	//19,10
	//19,2
	//16,2
	//16,0`

	//utils.Assert(70, part2(secondInput))

	ansPart2 := lazyPart2(input)
	goaocd.Submit(2, ansPart2, year, day)
	utils.Assert(ansPart2, part2(input))
}
