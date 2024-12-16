package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

const wall rune = '#'
const box rune = 'O'
const boxL rune = '['
const boxR rune = ']'
const robot rune = '@'
const up rune = '^'
const right rune = '>'
const down rune = 'v'
const left rune = '<'

var verticalMoves [][2]int

func moveBox(pos [2]int, move rune, grid [][]string) bool {
	newPos := walk(pos, move, grid)
	if newPos != pos {
		grid[pos[1]][pos[0]], grid[newPos[1]][newPos[0]] = grid[newPos[1]][newPos[0]], grid[pos[1]][pos[0]]
	}
	return newPos != pos
}

func moveLargeBoxVertically(pos [2]int, move rune, grid [][]string, char string) bool {
	var newLPos, newRPos [2]int
	if char == string(boxL) {
		newLPos = walk(pos, move, grid)
		newRPos = walk([2]int{pos[0] + 1, pos[1]}, move, grid)
	} else {
		newLPos = walk([2]int{pos[0] - 1, pos[1]}, move, grid)
		newRPos = walk(pos, move, grid)
	}

	if newLPos[1] != newRPos[1] {
		return false
	}

	moved := false

	if char == string(boxL) {
		grid[pos[1]][pos[0]], grid[newLPos[1]][newLPos[0]] = grid[newLPos[1]][newLPos[0]], grid[pos[1]][pos[0]]
		grid[pos[1]][pos[0]+1], grid[newRPos[1]][newRPos[0]] = grid[newRPos[1]][newRPos[0]], grid[pos[1]][pos[0]+1]
		moved = newLPos != pos && newRPos != [2]int{pos[0] + 1, pos[1]}
	} else {
		grid[pos[1]][pos[0]-1], grid[newLPos[1]][newLPos[0]] = grid[newLPos[1]][newLPos[0]], grid[pos[1]][pos[0]-1]
		grid[pos[1]][pos[0]], grid[newRPos[1]][newRPos[0]] = grid[newRPos[1]][newRPos[0]], grid[pos[1]][pos[0]]
		moved = newLPos != [2]int{pos[0] - 1, pos[1]} && newRPos != pos
	}

	if moved {
		verticalMoves = append(verticalMoves, newLPos, newRPos)
	}

	return moved
}

func resetVerticalMove(pos [2]int, move rune, grid [][]string) {
	x, y := pos[0], pos[1]
	switch move {
	case up:
		y++
	case down:
		y--
	}
	grid[y][x], grid[pos[1]][pos[0]] = grid[pos[1]][pos[0]], grid[y][x]
}

func walk(pos [2]int, move rune, grid [][]string) [2]int {
	x, y := pos[0], pos[1]
	switch move {
	case up:
		y--
	case right:
		x++
	case down:
		y++
	case left:
		x--
	}
	next := grid[y][x]
	if next == string(wall) {
		return pos
	}
	if next == string(box) && !moveBox([2]int{x, y}, move, grid) {
		return pos
	}
	if slices.Contains([]string{string(boxL), string(boxR)}, next) {
		if slices.Contains([]rune{right, left}, move) && !moveBox([2]int{x, y}, move, grid) {
			return pos
		}
		if slices.Contains([]rune{up, down}, move) && !moveLargeBoxVertically([2]int{x, y}, move, grid, next) {
			return pos
		}
	}
	return [2]int{x, y}
}

func solve(lines []string, moves string, chunky bool) (total int) {
	var pos [2]int
	grid := [][]string{}

	for y, line := range lines {
		grid = append(grid, []string{})
		for x, char := range line {
			if chunky && char == box {
				grid[y] = append(grid[y], string(boxL), string(boxR))
			} else if chunky && char == robot {
				grid[y] = append(grid[y], string(robot), ".")
				pos = [2]int{x * 2, y}
			} else if chunky {
				grid[y] = append(grid[y], string(char), string(char))
			} else {
				grid[y] = append(grid[y], string(char))
			}
			if char == robot && !chunky {
				pos = [2]int{x, y}
			}
		}
	}

	for _, move := range moves {
		verticalMoves = [][2]int{}
		newPos := walk(pos, move, grid)
		if newPos != pos {
			grid[pos[1]][pos[0]], grid[newPos[1]][newPos[0]] = grid[newPos[1]][newPos[0]], grid[pos[1]][pos[0]]
			pos = newPos
		} else {
			for i := len(verticalMoves) - 1; i >= 0; i-- {
				resetVerticalMove(verticalMoves[i], move, grid)
			}
		}
	}

	for y, row := range grid {
		for x, col := range row {
			if !chunky && col == string(box) {
				total += 100*y + x
			}

			if chunky && col == string(boxL) {
				total += 100*y + x
			}
		}
	}

	return total
}

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	lines := strings.Split(split[0], "\n")

	fmt.Println("Part 1:", solve(lines, split[1], false))
	fmt.Println("Part 2:", solve(lines, split[1], true))
}
