package main

import (
	"advent-of-code-2025/utils"
	"container/heap"
	"math"
	"strconv"
	"strings"

	"github.com/Olegas/goaocd"
	"github.com/draffensperger/golp"
)

type Machine struct {
	indicator     int
	length        int
	buttons       []int
	voltage       []int
	buttonsValues [][]int
	cols          int
}

func (machine Machine) BitIndicator(value int) string {
	var result []rune
	result = append(result, '[')
	for i := 0; i < machine.length; i++ {
		if value&(1<<i) != 0 {
			result = append(result, '#')
		} else {
			result = append(result, '.')
		}
	}
	result = append(result, ']')
	return string(result)
}

func parse(input string) []Machine {
	var machines []Machine
	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		parts := strings.Split(line, " ")
		indicatorText := parts[0][1 : len(parts[0])-1]
		indicator := 0
		for i, ch := range indicatorText {
			if ch == '#' {
				indicator |= 1 << i
			}
		}

		var buttons []int
		var intButtons = make([][]int, 0)
		for i := 1; i < len(parts)-1; i++ {
			buttonsText := parts[i][1 : len(parts[i])-1]
			buttonsValues := strings.Split(buttonsText, ",")
			buttonRows := make([]int, 0)
			for _, buttonValue := range buttonsValues {
				var value, _ = strconv.Atoi(buttonValue)
				buttonRows = append(buttonRows, value)
			}
			intButtons = append(intButtons, buttonRows)
			buttonMask := 0
			for _, b := range buttonsValues {
				var index, _ = strconv.Atoi(b)
				buttonMask |= 1 << index
			}
			buttons = append(buttons, buttonMask)
		}

		voltageText := parts[len(parts)-1]
		voltageText = voltageText[1 : len(voltageText)-1]
		voltageValues := strings.Split(voltageText, ",")
		voltage := make([]int, len(voltageValues))
		for i, v := range voltageValues {
			voltage[i], _ = strconv.Atoi(v)
		}

		machines = append(machines, Machine{indicator: indicator, buttons: buttons, voltage: voltage, length: len(indicatorText), buttonsValues: intButtons, cols: len(parts) - 2})
	}
	return machines
}

type Item struct {
	state  int
	clicks int
	arr    []int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].clicks < pq[j].clicks
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func part1(inputData string) int {
	var input = parse(inputData)
	minMoves := 0
	for _, machine := range input {
		startState := 0
		queue := make(PriorityQueue, 0)
		heap.Push(&queue, &Item{state: startState, clicks: 0})
		visited := make(map[int]bool)

		for len(queue) > 0 {
			current := heap.Pop(&queue).(*Item)
			println("Checking ", machine.BitIndicator(current.state), "with ", current.clicks, " clicks")

			if current.state == machine.indicator {
				println("Min clicks for ", machine.BitIndicator(machine.indicator), " is ", current.clicks)
				minMoves += current.clicks
				break
			}

			if visited[current.state] {
				continue
			}
			visited[current.state] = true

			for i, mask := range machine.buttons {
				var newState = current.state ^ mask
				if !visited[newState] {
					println("  Pushing ", i, " ", machine.BitIndicator(mask), " button on ", machine.BitIndicator(current.state), " with ", current.clicks+1, " clicks for", machine.BitIndicator(newState))
					heap.Push(&queue, &Item{state: newState, clicks: current.clicks + 1})
				}
			}

		}

	}
	return minMoves
}

func Hash(a []int) int {
	hash := 0
	for i := 0; i < len(a); i++ {
		hash = hash*31 + a[i]
	}
	return hash
}

func part2Slow(inputData string) int {
	var input = parse(inputData)
	minMoves := 0

	for machineIndex, machine := range input {
		queue := make(PriorityQueue, 0)
		startState := make([]int, machine.length)
		heap.Push(&queue, &Item{arr: startState, clicks: 0})
		visited := make(map[int]bool)
		var finalStateHash = Hash(machine.voltage)
		for len(queue) > 0 {
			current := heap.Pop(&queue).(*Item)

			var currentHash = Hash(current.arr)

			if currentHash == finalStateHash {
				//println("Min clicks for ", machine.BitIndicator(machine.indicator), " is ", current.clicks)
				minMoves += current.clicks
				break
			}

			if visited[currentHash] {
				continue
			}
			visited[currentHash] = true

			for _, mask := range machine.buttons {
				newState := make([]int, machine.length)
				copy(newState, current.arr)
				for bit := 0; bit < machine.length; bit++ {
					if mask&(1<<bit) != 0 {
						newState[bit] = current.arr[bit] + 1
					}
				}
				//println("  Pushing ", i, " ", machine.BitIndicator(mask), " button on ", current.arr, " with ", current.clicks+1, " clicks for", newState)
				heap.Push(&queue, &Item{arr: newState, clicks: current.clicks + 1})
			}
		}
		println(machineIndex, " / ", len(input))
	}
	return minMoves
}

func part2Fast(inputData string) int {
	var input = parse(inputData)
	clicks := 0

	for _, machine := range input {
		lp := golp.NewLP(machine.length, machine.cols)

		for row := range machine.length {
			var rowEntries []golp.Entry
			for col, values := range machine.buttonsValues {
				buttonPresent := false
				for i := 0; i < len(values); i++ {
					if values[i] == row {
						buttonPresent = true
						break
					}
				}
				if buttonPresent {
					rowEntries = append(rowEntries, golp.Entry{Col: col, Val: 1.0})
				}
			}

			_ = lp.AddConstraintSparse(rowEntries, golp.EQ, float64(machine.voltage[row]))
		}
		var objFn = make([]float64, machine.cols)
		for col := range machine.cols {
			objFn[col] = 1.0
			lp.SetInt(col, true)
		}
		lp.SetObjFn(objFn)
		lp.Solve()

		// additional check
		vars := lp.Variables()
		var sum = 0
		result := make([]int, machine.length)
		for col, buttons := range machine.buttonsValues {
			for _, b := range buttons {
				result[b] += int(vars[col])
			}
			sum += int(vars[col])
		}
		for i := range machine.length {
			if result[i] != machine.voltage[i] {
				println("Something is wrong ", i, result[i])
				panic("Something is wrong")
			}
		}

		clicks += int(math.Round(lp.Objective()))
	}
	return clicks
}

const day = 10
const year = 2025

func main() {
	var testInput = `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`

	var input = goaocd.Input(year, day)
	utils.Assert(7, part1(testInput))
	goaocd.Submit(1, part1(input), year, day)
	utils.Assert(33, part2Slow(testInput))
	println("Test for part 2 passed")
	utils.Assert(33, part2Fast(testInput))
	println("Test for fast part 2 passed")
	goaocd.Submit(2, part2Fast(input), year, day)
}
