package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"
)

func walk(guard [3]int, xObstacles map[int][]int, yObstacles map[int][]int, lines []string) ([][3]int, bool) {
	visited := [][3]int{guard}
	direction := "UP"

	for {
		if direction == "UP" {
			destination := -1
			if len(xObstacles[guard[0]]) > 0 {
				for _, v := range xObstacles[guard[0]] {
					if v >= guard[1] {
						break
					}
					destination = v
				}
			}
			for i := guard[1] - 1; i > destination; i-- {
				square := [3]int{guard[0], i, 0}
				if slices.Contains(visited, square) {
					return visited, true
				}
				visited = append(visited, [3]int{guard[0], i, 0})
			}
			if destination < 0 {
				break
			}
			guard[1] = destination + 1
			direction = "RIGHT"
		} else if direction == "RIGHT" {
			destination := len(lines[0])
			if len(yObstacles[guard[1]]) > 0 {
				for _, v := range yObstacles[guard[1]] {
					if v > guard[0] {
						destination = v
						break
					}
				}
			}
			for i := guard[0] + 1; i < destination; i++ {
				square := [3]int{i, guard[1], 1}
				if slices.Contains(visited, square) {
					return visited, true
				}
				visited = append(visited, [3]int{i, guard[1], 1})
			}
			if destination > len(lines[0])-1 {
				break
			}
			guard[0] = destination - 1
			direction = "DOWN"
		} else if direction == "DOWN" {
			destination := len(lines)
			if len(xObstacles[guard[0]]) > 0 {
				for _, v := range xObstacles[guard[0]] {
					if v > guard[1] {
						destination = v
						break
					}
				}
			}
			for i := guard[1] + 1; i < destination; i++ {
				square := [3]int{guard[0], i, 2}
				if slices.Contains(visited, square) {
					return visited, true
				}
				visited = append(visited, [3]int{guard[0], i, 2})
			}
			if destination > len(lines)-1 {
				break
			}
			guard[1] = destination - 1
			direction = "LEFT"
		} else if direction == "LEFT" {
			destination := -1
			if len(yObstacles[guard[1]]) > 0 {
				for _, v := range yObstacles[guard[1]] {
					if v >= guard[0] {
						break
					}
					destination = v
				}
			}
			for i := guard[0] - 1; i > destination; i-- {
				square := [3]int{i, guard[1], 3}
				if slices.Contains(visited, square) {
					return visited, true
				}
				visited = append(visited, [3]int{i, guard[1], 3})
			}
			if destination < 0 {
				break
			}
			guard[0] = destination + 1
			direction = "UP"
		}
	}

	return visited, false
}

func deepCloneMap(m map[int][]int) map[int][]int {
	clone := map[int][]int{}
	for k, v := range m {
		clone[k] = append([]int(nil), v...)
	}
	return clone
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var guard [3]int
	yObstacles := map[int][]int{}
	xObstacles := map[int][]int{}

	for y, line := range lines {
		for x, char := range line {
			if char == '^' {
				guard[0] = x
				guard[1] = y
				guard[2] = 0
			} else if char == '#' {
				xObstacles[x] = append(xObstacles[x], y)
				yObstacles[y] = append(yObstacles[y], x)
			}
		}
	}

	visited, _ := walk(guard, xObstacles, yObstacles, lines)
	unique := map[[2]int]bool{}
	for _, v := range visited {
		unique[[2]int{v[0], v[1]}] = true
	}
	fmt.Println("Part 1:", len(unique))

	var wg sync.WaitGroup

	count := 0
	for square := range unique {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if square[0] == guard[0] && square[1] == guard[1] {
				return
			}
			xObs := deepCloneMap(xObstacles)
			yObs := deepCloneMap(yObstacles)
			xObs[square[0]] = append(xObs[square[0]], square[1])
			yObs[square[1]] = append(yObs[square[1]], square[0])
			slices.Sort(xObs[square[0]])
			slices.Sort(yObs[square[1]])
			_, loops := walk(guard, xObs, yObs, lines)
			if loops {
				count++
			}
		}()
	}

	wg.Wait()

	fmt.Println("Part 2:", count)
}
