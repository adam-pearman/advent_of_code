package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var a, b, c int

func getComboOp(x int) int {
	switch x {
	case 0, 1, 2, 3:
		return x
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	default:
		return -1
	}
}

func adv(x int) {
	a /= int(math.Pow(2, float64(getComboOp(x))))
}

func bxl(x int) {
	b ^= x
}

func bst(x int) {
	b = getComboOp(x) % 8
}

func jnz() bool {
	return a != 0
}

func bxc() {
	b ^= c
}

func out(x int) string {
	return fmt.Sprint(getComboOp(x) % 8)
}

func bdv(x int) {
	b = a / int(math.Pow(2, float64(getComboOp(x))))
}

func cdv(x int) {
	c = a / int(math.Pow(2, float64(getComboOp(x))))
}

func run(instructions []string) []string {
	var output []string

	for i := 0; i < len(instructions); {
		instruction, _ := strconv.Atoi(instructions[i])
		x, _ := strconv.Atoi(instructions[i+1])
		switch instruction {
		case 0:
			adv(x)
		case 1:
			bxl(x)
		case 2:
			bst(x)
		case 3:
			if jnz() {
				i = x
				continue
			}
		case 4:
			bxc()
		case 5:
			output = append(output, out(x))
		case 6:
			bdv(x)
		case 7:
			cdv(x)
		}
		i += 2
	}

	return output
}

func main() {
	input, _ := os.ReadFile("input.txt")
	re := regexp.MustCompile(`\d+`)
	values := re.FindAllString(string(input), -1)

	a, _ = strconv.Atoi(values[0])
	b, _ = strconv.Atoi(values[1])
	c, _ = strconv.Atoi(values[2])

	instructions := values[3:]

	fmt.Println("Part 1:", strings.Join(run(instructions), ","))

	num := 0
	for i := 0; num == 0; i++ {
		a, b, c = i, 0, 0
		if strings.Join(run(instructions), ",") == "3,0" {
			num = i
			break
		}
	}

	for i := len(instructions) - 3; i >= 0; i-- {
		num *= 8
		a, b, c = num, 0, 0
		output := run(instructions)
		o, _ := strconv.Atoi(output[0])
		instruction, _ := strconv.Atoi(instructions[i])
		for o != instruction {
			num++
			a, b, c = num, 0, 0
			output = run(instructions)
			o, _ = strconv.Atoi(output[0])
		}
	}

	fmt.Println("Part 2:", num)
}
