package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var lines []string
var visited = [][2]int{}

func checkAbove(x, y int, char byte) bool {
	newY := y - 1
	return newY < 0 || lines[newY][x] != char
}

func checkRight(x, y int, char byte) bool {
	newX := x + 1
	return newX >= len(lines[0]) || lines[y][newX] != char
}

func checkBelow(x, y int, char byte) bool {
	newY := y + 1
	return newY >= len(lines) || lines[newY][x] != char
}

func checkLeft(x, y int, char byte) bool {
	newX := x - 1
	return newX < 0 || lines[y][newX] != char
}

func checkBelowRight(x, y int, char byte) bool {
	newX, newY := x+1, y+1
	if newX < len(lines[0]) && newY < len(lines) && lines[newY][newX] == char {
		return checkLeft(newX, newY, char) != checkAbove(newX, newY, char)
	}
	return false
}

func checkBelowLeft(x, y int, char byte) bool {
	newX, newY := x-1, y+1
	if newX >= 0 && newY < len(lines) && lines[newY][newX] == char {
		return checkRight(newX, newY, char) != checkAbove(newX, newY, char)
	}
	return false
}

func walk(x, y int, char byte) (int, int, int) {
	perim, area, corners := 0, 1, 0
	visited = append(visited, [2]int{x, y})

	top := checkAbove(x, y, char)
	right := checkRight(x, y, char)
	bottom := checkBelow(x, y, char)
	left := checkLeft(x, y, char)

	if top && right {
		corners++
	}
	if right && bottom {
		corners++
	}
	if bottom && left {
		corners++
	}
	if left && top {
		corners++
	}
	if checkBelowRight(x, y, char) {
		corners++
	}
	if checkBelowLeft(x, y, char) {
		corners++
	}

	if top {
		perim++
	} else if !slices.Contains(visited, [2]int{x, y - 1}) {
		p, a, c := walk(x, y-1, char)
		perim += p
		area += a
		corners += c
	}
	if right {
		perim++
	} else if !slices.Contains(visited, [2]int{x + 1, y}) {
		p, a, c := walk(x+1, y, char)
		perim += p
		area += a
		corners += c
	}
	if bottom {
		perim++
	} else if !slices.Contains(visited, [2]int{x, y + 1}) {
		p, a, c := walk(x, y+1, char)
		perim += p
		area += a
		corners += c
	}
	if left {
		perim++
	} else if !slices.Contains(visited, [2]int{x - 1, y}) {
		p, a, c := walk(x-1, y, char)
		perim += p
		area += a
		corners += c
	}

	return perim, area, corners
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines = strings.Split(strings.TrimSpace(string(input)), "\n")
	part1 := 0
	part2 := 0

	for y, line := range lines {
		for x, char := range line {
			if slices.Contains(visited, [2]int{x, y}) {
				continue
			}
			p, a, c := walk(x, y, byte(char))
			part1 += p * a
			part2 += a * c
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
