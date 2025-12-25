package main

import (
	"advent-of-code-2025/utils"
	"strings"

	"github.com/Olegas/goaocd"
)

func parse(input string) map[string][]string {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	graph := make(map[string][]string)
	for _, line := range lines {
		from := line[:3]
		to := strings.Split(line[5:], " ")
		graph[from] = to
	}
	return graph
}

func countPaths(start string, end string, graph map[string][]string) int {
	queue := []string{start}
	paths := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, neighbor := range graph[current] {
			queue = append(queue, neighbor)
			if neighbor == end {
				paths++
			}
		}
	}
	return paths
}

func part1(inputData string) int {
	var input = parse(inputData)
	return countPaths("you", "out", input)
}

func countPathsMemo(current string, end string, graph map[string][]string, memo map[string]int) int {
	if current == end {
		return 1
	}
	if _, ok := memo[current]; ok {
		return memo[current]
	}
	paths := 0
	for _, neighbor := range graph[current] {
		paths += countPathsMemo(neighbor, end, graph, memo)
	}
	memo[current] = paths
	return paths
}

type Node struct {
	name string
	dist int
}

func bfs(current string, end string, graph map[string][]string) int {
	queue := []Node{{current, 0}}
	visited := map[string]bool{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.name == end {
			return node.dist
		}
		if visited[node.name] {
			continue
		}
		for _, neighbor := range graph[node.name] {
			queue = append(queue, Node{neighbor, node.dist + 1})
		}
		visited[node.name] = true
	}
	panic("no solution found")
}
func part2(inputData string) int {
	var input = parse(inputData)
	distToFFT := bfs("svr", "fft", input)
	distToDAC := bfs("svr", "dac", input)
	var order []string
	if distToFFT < distToDAC {
		order = []string{"svr", "fft", "dac", "out"}
	} else {
		order = []string{"svr", "dac", "fft", "out"}
	}
	svrPaths := countPathsMemo(order[0], order[1], input, make(map[string]int))
	fftPaths := countPathsMemo(order[1], order[2], input, make(map[string]int))
	outPaths := countPathsMemo(order[2], order[3], input, make(map[string]int))
	return svrPaths * fftPaths * outPaths
}

const day = 11
const year = 2025

func main() {
	var testInput = `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`
	var input = goaocd.Input(year, day)
	utils.Assert(5, part1(testInput))
	//goaocd.Submit(1, part1(input), year, day)
	var secondInput = `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`
	utils.Assert(2, part2(secondInput))
	goaocd.Submit(2, part2(input), year, day)
}
