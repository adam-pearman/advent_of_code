package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dpad = map[string]map[string]string{
	"A": {"A": "A", "<": "v<<A", "^": "<A", ">": "vA", "v": "<vA"},
	"<": {"A": ">>^A", "<": "A", "^": ">^A", ">": ">>A", "v": ">A"},
	"^": {"A": ">A", "<": "v<A", "^": "A", ">": "v>A", "v": "vA"},
	">": {"A": "^A", "<": "<<A", "^": "<^A", ">": "A", "v": "<A"},
	"v": {"A": "^>A", "<": "<A", "^": "^A", ">": ">A", "v": "A"},
}

var numpad = map[string]map[string]string{
	"A": {"A": "A", "0": "<A", "1": "^<<A", "2": "<^A", "3": "^A", "4": "^^<<A", "5": "<^^A", "6": "^^A", "7": "^^^<<A", "8": "<^^^A", "9": "^^^A"},
	"0": {"A": ">A", "0": "A", "1": "^<A", "2": "^A", "3": "^>A", "4": "^^<A", "5": "^^A", "6": "^^>A", "7": "^^^<A", "8": "^^^A", "9": "^^^>A"},
	"1": {"A": ">>vA", "0": ">vA", "1": "A", "2": ">A", "3": ">>A", "4": "^A", "5": "^>A", "6": "^>>A", "7": "^^A", "8": "^^>A", "9": "^^>>A"},
	"2": {"A": "v>A", "0": "vA", "1": "<A", "2": "A", "3": ">A", "4": "<^A", "5": "^A", "6": "^>A", "7": "<^^A", "8": "^^A", "9": "^^>A"},
	"3": {"A": "vA", "0": "<vA", "1": "<<A", "2": "<A", "3": "A", "4": "<<^A", "5": "<^A", "6": "^A", "7": "<<^^A", "8": "<^^A", "9": "^^A"},
	"4": {"A": ">>vvA", "0": ">vvA", "1": "vA", "2": "v>A", "3": "v>>A", "4": "A", "5": ">A", "6": ">>A", "7": "^A", "8": "^>A", "9": "^>>A"},
	"5": {"A": "vv>A", "0": "vvA", "1": "<vA", "2": "vA", "3": "v>A", "4": "<A", "5": "A", "6": ">A", "7": "<^A", "8": "^A", "9": "^>A"},
	"6": {"A": "vvA", "0": "<vvA", "1": "<<vA", "2": "<vA", "3": "vA", "4": "<<A", "5": "<A", "6": "A", "7": "<<^A", "8": "<^A", "9": "^A"},
	"7": {"A": ">>vvvA", "0": ">vvvA", "1": "vvA", "2": "vv>A", "3": "vv>>A", "4": "vA", "5": "v>A", "6": "v>>A", "7": "A", "8": ">A", "9": ">>A"},
	"8": {"A": "vvv>A", "0": "vvvA", "1": "<vvA", "2": "vvA", "3": "vv>A", "4": "<vA", "5": "vA", "6": "v>A", "7": "<A", "8": "A", "9": ">A"},
	"9": {"A": "vvvA", "0": "<vvvA", "1": "<<vvA", "2": "<vvA", "3": "vvA", "4": "<<vA", "5": "<vA", "6": "vA", "7": "<<A", "8": "<A", "9": "A"},
}

var codes []string

var memo = make(map[string]int)

func lengthOfPair(i, j string, depth int) int {
	key := fmt.Sprintf("%s|%s|%d", i, j, depth)
	if value, ok := memo[key]; ok {
		return value
	}
	if depth == 1 {
		memo[key] = len(dpad[i][j])
		return memo[key]
	}
	result := 0
	pairs := getPairs("A" + dpad[i][j])
	for _, pair := range pairs {
		result += lengthOfPair(pair[0], pair[1], depth-1)
	}
	memo[key] = result
	return result
}

func getPairs(s string) [][2]string {
	pairs := make([][2]string, 0)
	for i := 0; i < len(s)-1; i++ {
		pairs = append(pairs, [2]string{s[i : i+1], s[i+1 : i+2]})
	}
	return pairs
}

func solve(depth int) int {
	total := 0
	for _, code := range codes {
		pairs := getPairs("A" + code)
		codeTotal := 0
		for _, pair := range pairs {
			sum := 0
			innerPairs := getPairs("A" + numpad[pair[0]][pair[1]])
			for _, innerPair := range innerPairs {
				sum += lengthOfPair(innerPair[0], innerPair[1], depth)
			}
			multiplier, _ := strconv.Atoi(code[:len(code)-1])
			codeTotal += multiplier * sum
		}
		total += codeTotal
	}
	return total
}

func main() {
	input, _ := os.ReadFile("input.txt")
	codes = strings.Split(strings.TrimSpace(string(input)), "\n")

	fmt.Println("Part 1:", solve(2))
	fmt.Println("Part 2:", solve(25))
}
