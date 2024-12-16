package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type direction struct {
	name string
	x, y int
}

var directions []direction = []direction{
	{"N", 0, -1},
	{"E", 1, 0},
	{"S", 0, 1},
	{"W", -1, 0},
}

func heuristic(pos, goal [2]int) int {
	return int(math.Abs(float64(pos[0]-goal[0])) + math.Abs(float64(pos[1]-goal[1])))
}

func directionIndex(name string) int {
	for i, d := range directions {
		if d.name == name {
			return i
		}
	}
	return -1
}

func getDirections(index int) []int {
	if index == 0 {
		return []int{0, 1, 3}
	}
	if index == 1 {
		return []int{0, 1, 2}
	}
	if index == 2 {
		return []int{1, 2, 3}
	}
	if index == 3 {
		return []int{0, 2, 3}
	}
	return []int{}
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var start, end [2]int
	for y, line := range lines {
		if strings.Contains(line, "S") {
			start = [2]int{strings.Index(line, "S"), y}
		}
		if strings.Contains(line, "E") {
			end = [2]int{strings.Index(line, "E"), y}
		}
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	for i := range directions {
		cost := 1000
		if i == 1 {
			cost = 0
		}
		heap.Push(pq, &Item{pos: start, dirIndex: i, priority: 0, cost: cost, path: [][2]int{start}})
	}

	visited := make(map[[3]int]int)
	minScore := math.MaxInt
	paths := make([][2]int, 0)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		state := [3]int{current.pos[0], current.pos[1], current.dirIndex}

		if minCost, ok := visited[state]; ok && current.cost > minCost {
			continue
		}
		visited[state] = current.cost

		if current.pos == end {
			if current.cost < minScore {
				minScore = current.cost
				paths = append(paths, current.path...)
			}
			if current.cost == minScore {
				paths = append(paths, current.path...)
			}
			continue
		}

		for i, d := range directions {
			newPos := [2]int{current.pos[0] + d.x, current.pos[1] + d.y}
			newChar := lines[newPos[1]][newPos[0]]
			if newChar == '#' {
				continue
			}
			moveCost := 1
			if i != current.dirIndex {
				moveCost += 1000
			}
			newCost := current.cost + moveCost
			priority := newCost + heuristic(newPos, end)

			newPath := append(slices.Clone(current.path), newPos)

			heap.Push(pq, &Item{pos: newPos, dirIndex: i, priority: priority, cost: newCost, path: newPath})
		}
	}

	fmt.Println("Part 1:", minScore)

	unique := make(map[[2]int]bool)
	for _, p := range paths {
		unique[p] = true
	}
	fmt.Println("Part 2:", len(unique))
}
