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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(lSlice, rSlice []string) {
	lSlice = slices.Clone(lSlice)
	rSlice = slices.Clone(rSlice)

	slices.Sort(lSlice)
	slices.Sort(rSlice)

	total := 0

	for i, l := range lSlice {
		l, err := strconv.Atoi(l)
		check(err)
		r, err := strconv.Atoi(rSlice[i])
		check(err)

		total += int(math.Abs(float64(l) - float64(r)))
	}

	fmt.Println("Part 1:", total)
}

func part2(lSlice, rSlice []string) {
	total := 0
	for _, l := range lSlice {
		sum := 0
		for _, r := range rSlice {
			if l == r {
				sum++
			}
		}
		l, err := strconv.Atoi(l)
		check(err)
		total += sum * l
	}

	fmt.Println("Part 2:", total)
}

func main() {
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	lSlice := []string{}
	rSlice := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		lSlice = append(lSlice, split[0])
		rSlice = append(rSlice, split[len(split)-1])
	}

	part1(lSlice, rSlice)
	part2(lSlice, rSlice)
}
