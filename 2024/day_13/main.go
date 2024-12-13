package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func solveSimEq(x1, y1, c1, x2, y2, c2 int) (x int, y int) {
	det, detX, detY := x1*y2-y1*x2, c1*y2-y1*c2, x1*c2-c1*x2
	x, y = detX/det, detY/det
	if x1*x+y1*y == c1 && x2*x+y2*y == c2 {
		return x, y
	}
	return 0, 0
}

type machine struct {
	a1, b1, p1, a2, b2, p2 int
}

func newMachine(values [6]string) machine {
	a1, _ := strconv.Atoi(values[0])
	a2, _ := strconv.Atoi(values[1])
	b1, _ := strconv.Atoi(values[2])
	b2, _ := strconv.Atoi(values[3])
	p1, _ := strconv.Atoi(values[4])
	p2, _ := strconv.Atoi(values[5])

	return machine{a1, b1, p1, a2, b2, p2}
}

func (m *machine) solve(offset int) int {
	p1 := m.p1 + offset
	p2 := m.p2 + offset
	a, b := solveSimEq(m.a1, m.b1, p1, m.a2, m.b2, p2)
	return a*3 + b
}

func main() {
	input, _ := os.ReadFile("input.txt")
	machines := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	part1 := 0
	part2 := 0
	re := regexp.MustCompile(`(\d+)`)

	for _, machine := range machines {
		v := re.FindAllString(machine, -1)
		m := newMachine([6]string{v[0], v[1], v[2], v[3], v[4], v[5]})
		part1 += m.solve(0)
		part2 += m.solve(10000000000000)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
