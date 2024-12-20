package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Position struct {
	x, y int
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var start, end Position
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				start = Position{x, y}
			} else if c == 'E' {
				end = Position{x, y}
			}
		}
	}

	history := []Position{start}
	visited := make(map[Position]bool)
	visited[start] = true

	for history[len(history)-1] != end {
		current := history[len(history)-1]

		neighbors := []Position{
			{current.x, current.y - 1},
			{current.x, current.y + 1},
			{current.x - 1, current.y},
			{current.x + 1, current.y},
		}
		validNeighbors := []Position{}
		for _, n := range neighbors {
			if n.x >= 0 && n.x < len(lines[0]) && n.y >= 0 && n.y < len(lines) && lines[n.y][n.x] != '#' && !visited[n] {
				validNeighbors = append(validNeighbors, n)
			}
		}

		if len(validNeighbors) == 0 {
			break
		}

		next := validNeighbors[0]
		if len(history) >= 2 {
			prev := history[len(history)-2]
			for _, n := range validNeighbors {
				if n != prev {
					next = n
					break
				}
			}
		}

		visited[next] = true
		history = append(history, next)
	}

	part1 := 0
	part2 := 0
	threshold := 100
	for n, pos1 := range history {
		if n+threshold+2 >= len(history) {
			break
		}
		for m := 0; m < len(history[n+threshold+2:]); m++ {
			pos2 := history[n+m+threshold+2]
			distance := math.Abs(float64(pos1.x-pos2.x)) + math.Abs(float64(pos1.y-pos2.y))

			if distance == 2 {
				part1++
			}

			if distance <= math.Min(20, float64(m+2)) {
				part2++
			}
		}
	}

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
