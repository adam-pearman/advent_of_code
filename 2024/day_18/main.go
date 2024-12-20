package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func heuristic(a, b [2]int) int {
	return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}

func findPath(grid [][space]int, start, end [2]int) [][2]int {
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	pq := &PriorityQueue{}
	heap.Init(pq)

	gScore := make(map[[2]int]int)
	parents := make(map[[2]int][2]int)

	heap.Push(pq, &Node{Position: start, Cost: 0})
	gScore[start] = 0

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node).Position

		if current == end {
			var path [][2]int
			for current != start {
				path = append(path, current)
				current = parents[current]
			}
			return path
		}

		for _, direction := range directions {
			neighbor := [2]int{current[0] + direction[0], current[1] + direction[1]}

			if neighbor[0] < 0 || neighbor[0] >= len(grid) || neighbor[1] < 0 || neighbor[1] >= len(grid[0]) || grid[neighbor[0]][neighbor[1]] == 1 {
				continue
			}

			tentativeG := gScore[current] + 1

			if g, ok := gScore[neighbor]; !ok || tentativeG < g {
				gScore[neighbor] = tentativeG
				fScore := tentativeG + heuristic(neighbor, end)
				heap.Push(pq, &Node{Position: neighbor, Cost: fScore})
				parents[neighbor] = current
			}
		}
	}

	return nil
}

const space = 71
const bytes = 1024

func main() {
	input, _ := os.ReadFile("input.txt")
	obstacles := strings.Split(strings.TrimSpace(string(input)), "\n")
	grid := make([][space]int, space)
	start := [2]int{0, 0}
	end := [2]int{space - 1, space - 1}

	for i := 0; i < bytes; i++ {
		obstacle := strings.Split(obstacles[i], ",")
		x, _ := strconv.Atoi(obstacle[0])
		y, _ := strconv.Atoi(obstacle[1])
		grid[y][x] = 1
	}

	path := findPath(grid, start, end)
	fmt.Println("Part 1:", len(path))

	for i := bytes; i < len(obstacles); i++ {
		obstacle := strings.Split(obstacles[i], ",")
		x, _ := strconv.Atoi(obstacle[0])
		y, _ := strconv.Atoi(obstacle[1])
		grid[y][x] = 1
		path = findPath(grid, start, end)
		if path == nil {
			fmt.Println("Part 2:", obstacles[i])
			break
		}
	}
}
