package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var rules = map[string][]string{}
var lines [][]string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	rulesDefined := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			rulesDefined = true
		} else if rulesDefined {
			lines = append(lines, strings.Split(line, ","))
		} else {
			rule := strings.Split(line, "|")
			rules[rule[0]] = append(rules[rule[0]], rule[1])
		}
	}

	correctTotal := 0
	incorrectTotal := 0
	for _, line := range lines {
		pass := true
		for i, page := range line {
			for _, allowed := range line[i+1:] {
				if !slices.Contains(rules[page], string(allowed)) {
					pass = false
					break
				}
			}
			if !pass {
				break
			}
		}
		if pass {
			mid, err := strconv.Atoi(line[len(line)/2])
			check(err)
			correctTotal += mid
		} else {
			for i := 0; i < len(line); {
				swapped := false
				for j, allowed := range line[i+1:] {
					if !slices.Contains(rules[line[i]], string(allowed)) {
						line[i], line[j+i+1] = line[j+i+1], line[i]
						swapped = true
						break
					}
				}
				if !swapped {
					i++
				}
			}
			mid, err := strconv.Atoi(line[len(line)/2])
			check(err)
			incorrectTotal += mid
		}
	}

	fmt.Println("Part 1: ", correctTotal)
	fmt.Println("Part 2: ", incorrectTotal)
}
