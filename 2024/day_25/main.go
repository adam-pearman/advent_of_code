package main

import (
	"fmt"
	"os"
	"strings"
)

func getLockNums(schematic string) [5]int {
	nums := [5]int{}
	for i, line := range strings.Split(schematic, "\n") {
		for j, char := range line {
			if char == '#' {
				nums[j] = i
			}
		}
	}
	return nums
}

func getKeyNums(schematic string) [5]int {
	nums := [5]int{}
	for i, line := range strings.Split(schematic, "\n") {
		for j, char := range line {
			if char == '#' && nums[j] == 0 {
				nums[j] = 6 - i
			}
		}
	}
	return nums
}

func main() {
	input, _ := os.ReadFile("input.txt")
	locks := [][5]int{}
	keys := [][5]int{}

	for _, schematic := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		if strings.HasPrefix(schematic, "#") {
			locks = append(locks, getLockNums(schematic))
		} else {
			keys = append(keys, getKeyNums(schematic))
		}
	}

	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			fits := true
			for i := 0; i < 5; i++ {
				if lock[i]+key[i] > 5 {
					fits = false
					break
				}
			}
			if fits {
				count++
			}
		}
	}

	fmt.Println("Part 1:", count)
}
