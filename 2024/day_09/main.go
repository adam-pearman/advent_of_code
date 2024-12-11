package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(input string) {
	blocks := []string{}

	for i, char := range input {
		count, _ := strconv.Atoi(string(char))
		if i%2 == 0 {
			for j := 0; j < count; j++ {
				blocks = append(blocks, strconv.Itoa(i/2))
			}
		} else {
			blocks = append(blocks, strings.Split(strings.Repeat(".", count), "")...)
		}
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i] == "." {
			continue
		}
		min := 0
		for j := min; j < i; j++ {
			if blocks[j] == "." {
				blocks[j], blocks[i] = blocks[i], blocks[j]
				min = j + 1
				break
			}
		}
	}

	checksum := 0

	for i, block := range blocks {
		if block == "." {
			break
		}
		value, _ := strconv.Atoi(block)
		checksum += value * i
	}

	fmt.Println("Part 1:", checksum)
}

func part2(input string) {
	blocks := [][]string{}
	spaces := []int{}

	for i, char := range input {
		if char == '0' {
			continue
		}
		count, _ := strconv.Atoi(string(char))
		value := "."
		if i%2 == 0 {
			value = strconv.Itoa(i / 2)
		} else {
			spaces = append(spaces, len(blocks))
		}
		blocks = append(blocks, []string{value})
		for j := 1; j < count; j++ {
			blocks[len(blocks)-1] = append(blocks[len(blocks)-1], value)
		}
	}

	visited := []string{"."}

	for i := len(blocks) - 1; i >= 0; i-- {
		if slices.Contains(visited, blocks[i][0]) {
			continue
		}
		visited = append(visited, blocks[i][0])
		for j, space := range spaces {
			if space >= i {
				break
			}
			if len(blocks[space]) >= len(blocks[i]) {
				diff := len(blocks[space]) - len(blocks[i])
				blocks = append(blocks[:space], append([][]string{blocks[space][diff:]}, blocks[space:]...)...)
				space++
				blocks[space] = blocks[space][:diff]
				blocks[space-1], blocks[i+1] = blocks[i+1], blocks[space-1]
				spaces = append(spaces, i)
				slices.Sort(spaces)
				if len(blocks[space]) == 0 {
					blocks = append(blocks[:space], blocks[space+1:]...)
					spaces = append(spaces[:j], spaces[j+1:]...)
				} else {
					for k := j; k < len(spaces); k++ {
						spaces[k]++
					}
				}
				break
			}
		}
	}

	checksum := 0
	multiplier := 0

	for _, block := range blocks {
		id := block[0]
		if id == "." {
			multiplier += len(block)
			continue
		}
		for _, char := range block {
			value, _ := strconv.Atoi(char)
			checksum += value * multiplier
			multiplier++
		}
	}

	fmt.Println("Part 2:", checksum)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	trimmed := strings.TrimSpace(string(input))

	part1(trimmed)
	part2(trimmed)
}
