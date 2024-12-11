package main

import (
	"fmt"
	"os"
	"strings"
)

func getAntinodes(freqs map[string][][2]int, lines []string, part2 bool) int {
	antinodes := [][2]int{}

	for _, freq := range freqs {
		for i, antenna1 := range freq {
			for j := i + 1; j < len(freq); j++ {
				antenna2 := freq[j]
				if part2 {
					antinodes = append(antinodes, antenna1, antenna2)
				}
				xDiff := antenna1[0] - antenna2[0]
				yDiff := antenna1[1] - antenna2[1]
				antinode1 := [2]int{antenna1[0] + xDiff, antenna1[1] + yDiff}
				antinode2 := [2]int{antenna2[0] - xDiff, antenna2[1] - yDiff}
				antinodes = append(antinodes, antinode1, antinode2)
				for part2 {
					newAntinode := false
					if antinode1[0] >= 0 && antinode1[1] >= 0 && antinode1[0] < len(lines[0]) && antinode1[1] < len(lines) {
						antinode1 = [2]int{antinode1[0] + xDiff, antinode1[1] + yDiff}
						antinodes = append(antinodes, antinode1)
						newAntinode = true
					}
					if antinode2[0] >= 0 && antinode2[1] >= 0 && antinode2[0] < len(lines[0]) && antinode2[1] < len(lines) {
						antinode2 = [2]int{antinode2[0] - xDiff, antinode2[1] - yDiff}
						antinodes = append(antinodes, antinode2)
						newAntinode = true
					}
					if !newAntinode {
						break
					}
				}
			}
		}
	}

	unique := map[[2]int]bool{}
	for _, v := range antinodes {
		if v[0] < 0 || v[1] < 0 || v[0] >= len(lines[0]) || v[1] >= len(lines) {
			continue
		}
		unique[v] = true
	}

	return len(unique)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	freqs := map[string][][2]int{}

	for y, line := range lines {
		split := strings.Split(line, "")
		for x, char := range split {
			if char != "." {
				freqs[char] = append(freqs[char], [2]int{x, y})
			}
		}
	}

	fmt.Println("Part 1: ", getAntinodes(freqs, lines, false))
	fmt.Println("Part 2: ", getAntinodes(freqs, lines, true))
}
