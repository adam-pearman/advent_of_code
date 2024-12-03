package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

var asc = "+"
var desc = "-"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isSafe(line []string) bool {
	var dir *string
	safe := true

	for i := 0; i < len(line)-1; i++ {
		curr, err := strconv.Atoi(line[i])
		check(err)
		next, err := strconv.Atoi(line[i+1])
		check(err)

		diff := math.Abs(float64(curr - next))

		if diff < 1 || diff > 3 {
			safe = false
			break
		}

		if dir == nil && curr < next {
			dir = &asc
		} else if dir == nil {
			dir = &desc
		}

		if dir == &asc && curr > next {
			safe = false
			break
		}

		if dir == &desc && curr < next {
			safe = false
			break
		}
	}

	return safe
}

func isSafeRecursive(line []string, faults int) bool {
	var dir *string

	faultsCopy := faults

	for i := 0; i < len(line)-1; i++ {
		curr, err := strconv.Atoi(line[i])
		check(err)
		next, err := strconv.Atoi(line[i+1])
		check(err)

		diff := math.Abs(float64(curr - next))

		if dir == nil && curr < next {
			dir = &asc
		} else if dir == nil {
			dir = &desc
		}

		if diff < 1 || diff > 3 {
			faults++
		} else if dir == &asc && curr > next {
			faults++
		} else if dir == &desc && curr < next {
			faults++
		}

		if faults > 1 {
			return false
		}

		if faultsCopy < faults {
			skipPrev := false
			skipCurr := false
			skipNext := false

			if i > 0 {
				linePrev := slices.Clone(line)
				linePrev = append(linePrev[:i-1], linePrev[i:]...)
				skipPrev = isSafeRecursive(linePrev, faults)
			}
			if i < len(line)-1 {
				lineCurr := slices.Clone(line)
				lineCurr = append(lineCurr[:i], lineCurr[i+1:]...)
				skipCurr = isSafeRecursive(lineCurr, faults)
			}
			if i < len(line)-2 {
				lineNext := slices.Clone(line)
				lineNext = append(lineNext[:i+1], lineNext[i+2:]...)
				skipNext = isSafeRecursive(lineNext, faults)
			} else {
				skipNext = isSafeRecursive(line[:i+1], faults)
			}

			if !skipPrev && !skipCurr && !skipNext {
				return false
			}

			return true
		}
	}

	return true
}

func part1(lines [][]string) {
	safeCount := 0

	for _, line := range lines {
		safe := isSafe(line)
		if safe {
			safeCount++
		}
	}

	fmt.Println("Part 1:", safeCount)
}

func part2(lines [][]string) {
	safeCount := 0

	for _, line := range lines {
		safe := isSafeRecursive(line, 0)
		if safe {
			safeCount++
		}
	}

	fmt.Println("Part 2:", safeCount)
}

func main() {
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	lines := [][]string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.Split(line, " "))
	}

	part1(lines)
	part2(lines)
}
