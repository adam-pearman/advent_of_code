package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func multDir(dir []byte) int {
	re := regexp.MustCompile(`\d+`)
	nums := re.FindAll(dir, -1)
	digits := []int{}
	for _, num := range nums {
		i, err := strconv.Atoi(string(num))
		check(err)
		digits = append(digits, i)
	}
	return digits[0] * digits[1]
}

func part1(b strings.Builder) {
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	dirs := r.FindAll([]byte(b.String()), -1)
	total := 0

	for _, dir := range dirs {
		total = total + multDir(dir)
	}

	fmt.Println("Part 1:", total)
}

func part2(b strings.Builder) {
	r := regexp.MustCompile(`do\(\)|don\'t\(\)|mul\(\d+,\d+\)`)
	dirs := r.FindAll([]byte(b.String()), -1)
	total := 0
	do := true

	for _, dir := range dirs {
		if string(dir) == "don't()" {
			do = false
			continue
		}

		if string(dir) == "do()" {
			do = true
			continue
		}

		if do {
			total = total + multDir(dir)
		}
	}

	fmt.Println("Part 2:", total)
}

func main() {
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	var b strings.Builder
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		b.WriteString(line)
	}

	part1(b)
	part2(b)
}
