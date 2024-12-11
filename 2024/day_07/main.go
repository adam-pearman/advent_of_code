package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func isPossible(result int, values []string, part2 bool) bool {
	curr, _ := strconv.Atoi(values[0])
	if len(values) < 2 {
		return result == curr
	}
	next, _ := strconv.Atoi(values[1])
	sum := curr + next
	product := curr * next
	concat, _ := strconv.Atoi(strings.Join(values[:2], ""))
	canSum := false
	canMult := false
	canConcat := false
	if sum <= result {
		sumValues := slices.Clone(values[2:])
		sumValues = append([]string{strconv.Itoa(sum)}, sumValues...)
		canSum = isPossible(result, sumValues, part2)
	}
	if product <= result {
		multValues := slices.Clone(values[2:])
		multValues = append([]string{strconv.Itoa(product)}, multValues...)
		canMult = isPossible(result, multValues, part2)
	}
	if concat <= result && part2 {
		concatValues := slices.Clone(values[2:])
		concatValues = append([]string{strconv.Itoa(concat)}, concatValues...)
		canConcat = isPossible(result, concatValues, part2)
	}

	return canSum || canMult || canConcat
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	p1Total := 0
	p2Total := 0

	var wg sync.WaitGroup

	for _, line := range lines {
		wg.Add(1)

		go func(line string) {
			defer wg.Done()
			split := strings.Split(line, ": ")
			result, _ := strconv.Atoi(split[0])
			values := strings.Split(split[1], " ")

			if isPossible(result, values, false) {
				p1Total += result
			}

			if isPossible(result, values, true) {
				p2Total += result
			}
		}(line)
	}

	wg.Wait()

	fmt.Println("Part 1: ", p1Total)
	fmt.Println("Part 2: ", p2Total)
}
