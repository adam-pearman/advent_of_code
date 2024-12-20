package main

import (
	"fmt"
	"os"
	"strings"
)

var memo = make(map[string]int)

func getCount(pattern string, towels []string) int {
	if _, ok := memo[pattern]; ok {
		return memo[pattern]
	}
	count := 0
	for _, towel := range towels {
		if strings.HasPrefix(pattern, towel) {
			count += getCount(pattern[len(towel):], towels)
		}
	}
	if pattern == "" {
		count++
	}
	memo[pattern] = count
	return count
}

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	towels := strings.Split(split[0], ", ")
	patterns := strings.Split(split[1], "\n")

	counts := make([]int, 0)
	for _, pattern := range patterns {
		counts = append(counts, getCount(pattern, towels))
	}

	part1 := 0
	part2 := 0
	for _, count := range counts {
		if count > 0 {
			part1++
			part2 += count
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
