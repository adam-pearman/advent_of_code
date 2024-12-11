package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lines []string
var moves = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
var trailheads = [][2]int{}

func walk(char string, x, y int) {
	if char == "0" {
		trailheads = append(trailheads, [2]int{x, y})
		return
	}

	for _, move := range moves {
		newX, newY := x+move[0], y+move[1]
		if newX < 0 || newX >= len(lines[0]) || newY < 0 || newY >= len(lines) {
			continue
		}
		curr, _ := strconv.Atoi(char)
		tar, _ := strconv.Atoi(string(lines[newY][newX]))
		if curr-tar != 1 {
			continue
		}
		walk(fmt.Sprint(tar), newX, newY)
	}
}

func part1() {
	count := 0

	for y, line := range lines {
		for x, char := range line {
			trailheads = [][2]int{}
			if char == '9' {
				walk(string(char), x, y)
				unique := map[[2]int]bool{}
				for _, t := range trailheads {
					unique[t] = true
				}
				count += len(unique)
			}
		}
	}

	fmt.Println("Part 1:", count)
}

func part2() {
	trailheads = [][2]int{}

	for y, line := range lines {
		for x, char := range line {
			if char == '9' {
				walk(string(char), x, y)
			}
		}
	}

	fmt.Println("Part 2:", len(trailheads))
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines = strings.Split(strings.TrimSpace(string(input)), "\n")

	part1()
	part2()
}
