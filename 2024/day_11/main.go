package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input []byte
var stones []int
var arraySize int

func blink(stone int) []int {
	output := []int{}

	if stone == 0 {
		output = append(output, 1)
	} else if len(fmt.Sprint(stone))%2 == 0 {
		stoneStr := fmt.Sprint(stone)
		mid := len(stoneStr) / 2
		first, _ := strconv.Atoi(stoneStr[:mid])
		second, _ := strconv.Atoi(stoneStr[mid:])
		output = append(output, first, second)
	} else {
		output = append(output, stone*2024)
	}

	return output
}

func countStones(stone, n int, cache map[int][]int) int {
	if _, ok := cache[stone]; ok {
		if cache[stone][n-1] != 0 {
			return cache[stone][n-1]
		}
	} else {
		cache[stone] = make([]int, arraySize)
	}

	if n == 1 {
		cache[stone][n-1] = len(blink(stone))
		return len(blink(stone))
	}

	sum := 0

	for _, blinked := range blink(stone) {
		sum += countStones(blinked, n-1, cache)
	}

	cache[stone][n-1] = sum
	return sum
}

func run(n int) int {
	cache := make(map[int][]int)
	sum := 0
	arraySize = n

	for _, stone := range stones {
		sum += countStones(stone, n, cache)
	}

	return sum
}

func main() {
	input, _ = os.ReadFile("input.txt")
	for _, stone := range strings.Split(strings.TrimSpace(string(input)), " ") {
		num, _ := strconv.Atoi(stone)
		stones = append(stones, num)
	}

	fmt.Println("Part 1:", run(25))
	fmt.Println("Part 2:", run(75))
}
