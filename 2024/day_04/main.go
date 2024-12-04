package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

var xLen int
var yLen int
var lines [][]string

const search string = "XMAS"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSurroundingCoords(coord Coordinate) []Coordinate {
	coords := []Coordinate{}

	if coord.Y > 0 {
		coords = append(coords, coord.GetUp())
	}
	if coord.X < xLen-1 && coord.Y > 0 {
		coords = append(coords, coord.GetUpRight())
	}
	if coord.X < xLen-1 {
		coords = append(coords, coord.GetRight())
	}
	if coord.X < xLen-1 && coord.Y < yLen-1 {
		coords = append(coords, coord.GetDownRight())
	}
	if coord.Y < yLen-1 {
		coords = append(coords, coord.GetDown())
	}
	if coord.X > 0 && coord.Y < yLen-1 {
		coords = append(coords, coord.GetDownLeft())
	}
	if coord.X > 0 {
		coords = append(coords, coord.GetLeft())
	}
	if coord.X > 0 && coord.Y > 0 {
		coords = append(coords, coord.GetUpLeft())
	}

	return coords
}

func walk(coord Coordinate, letterIndex int) bool {
	if letterIndex >= len(search) {
		return true
	}

	if coord.X < 0 || coord.Y < 0 || coord.X >= xLen || coord.Y >= yLen {
		return false
	}

	if lines[coord.Y][coord.X] == string(search[letterIndex]) {
		return walk(coord.Step(), letterIndex+1)
	}

	return false
}

func part1() {
	count := 0

	for y, line := range lines {
		for x, char := range line {
			if char == string(search[0]) {
				coords := getSurroundingCoords(Coordinate{x, y, ""})
				for _, coord := range coords {
					found := walk(coord, 1)
					if found {
						count++
					}
				}
			}
		}
	}

	fmt.Println("Part 1:", count)
}

func getCornerCoords(coord Coordinate) []Coordinate {
	coords := []Coordinate{}

	if coord.X < xLen-1 && coord.Y > 0 {
		coords = append(coords, coord.GetUpRight())
	}
	if coord.X < xLen-1 && coord.Y < yLen-1 {
		coords = append(coords, coord.GetDownRight())
	}
	if coord.X > 0 && coord.Y < yLen-1 {
		coords = append(coords, coord.GetDownLeft())
	}
	if coord.X > 0 && coord.Y > 0 {
		coords = append(coords, coord.GetUpLeft())
	}

	return coords
}

func checkCorners(coords []Coordinate) bool {
	corners := map[string]string{"UR": "", "DR": "", "DL": "", "UL": ""}

	for _, coord := range coords {
		allowed := []string{"M", "S"}
		if !slices.Contains(allowed, lines[coord.Y][coord.X]) {
			break
		}
		corners[coord.Direction] = lines[coord.Y][coord.X]
	}

	if slices.Contains(slices.Collect(maps.Values(corners)), "") {
		return false
	}

	if corners["UR"] == corners["DL"] || corners["DR"] == corners["UL"] {
		return false
	}

	return true
}

func part2() {
	count := 0

	for y, line := range lines {
		for x, char := range line {
			if char == "A" {
				coords := getCornerCoords(Coordinate{x, y, ""})
				if len(coords) < 4 {
					continue
				}
				if checkCorners(coords) {
					count++
				}
			}
		}
	}

	fmt.Println("Part 2:", count)
}

func main() {
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.Split(line, ""))
	}

	xLen = len(lines[0])
	yLen = len(lines)

	part1()
	part2()
}
