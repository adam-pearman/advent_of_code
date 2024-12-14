package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const gridWidth int = 101
const gridHeight int = 103
const seconds int = 100

var robots []robot
var lines []string

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

type robot struct {
	pX, pY, vX, vY int
}

func newRobot(pX, pY, vX, vY string) robot {
	return robot{toInt(pX), toInt(pY), toInt(vX), toInt(vY)}
}

func (r *robot) move() {
	r.pX += r.vX
	r.pY += r.vY

	if r.pX < 0 {
		r.pX += gridWidth
	}
	if r.pX >= gridWidth {
		r.pX %= gridWidth
	}

	if r.pY < 0 {
		r.pY += gridHeight
	}
	if r.pY >= gridHeight {
		r.pY %= gridHeight
	}
}

func part1() {
	bots := slices.Clone(robots)

	for i := 1; i <= seconds; i++ {
		for j, bot := range bots {
			bot.move()
			bots[j] = bot
		}
	}

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	for _, bot := range bots {
		if bot.pX < gridWidth/2 && bot.pY < gridHeight/2 {
			q1++
		} else if bot.pX > gridWidth/2 && bot.pY < gridHeight/2 {
			q2++
		} else if bot.pX < gridWidth/2 && bot.pY > gridHeight/2 {
			q3++
		} else if bot.pX > gridWidth/2 && bot.pY > gridHeight/2 {
			q4++
		}
	}

	fmt.Println("Part 1:", q1*q2*q3*q4)
}

func calculateVariance(values []int) int {
	n := len(values)
	if n == 0 {
		return 0
	}

	sum := 0
	for _, value := range values {
		sum += value
	}
	mean := sum / n

	sqDiff := 0
	for _, value := range values {
		diff := value - mean
		sqDiff += diff * diff
	}

	return sqDiff / n
}

func part2() {
	bots := slices.Clone(robots)
	variances := make([]int, gridHeight*gridWidth)

	for i := 0; i < gridHeight*gridWidth; i++ {
		xValues := make([]int, len(bots))
		yValues := make([]int, len(bots))
		for j, bot := range bots {
			if i > 0 {
				bot.move()
				bots[j] = bot
			}
			xValues[j] = bot.pX
			yValues[j] = bot.pY
		}

		variances[i] = calculateVariance(xValues) + calculateVariance(yValues)
	}

	fmt.Println("Part 2:", slices.Index(variances, slices.Min(variances)))
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines = strings.Split(strings.TrimSpace(string(input)), "\n")
	re := regexp.MustCompile(`-?\d+`)

	for _, line := range lines {
		values := re.FindAllString(line, -1)
		robot := newRobot(values[0], values[1], values[2], values[3])
		robots = append(robots, robot)
	}

	part1()
	part2()
}
