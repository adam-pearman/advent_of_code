package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Instruction struct {
	input1, input2, gate, output string
}

func solve(instructions []Instruction, values map[string]int) int {
	v := maps.Clone(values)
	outputs := []string{}
	visited := map[string]struct{}{}
	i := 0
	for len(visited) != len(instructions) {
		if i == len(instructions) {
			i = 0
		}
		instruction := instructions[i]
		i++
		_, input1 := v[instruction.input1]
		_, input2 := v[instruction.input2]
		_, seen := visited[instruction.output]
		if !input1 || !input2 || seen {
			continue
		}
		if instruction.gate == "AND" {
			v[instruction.output] = v[instruction.input1] & v[instruction.input2]
		} else if instruction.gate == "OR" {
			v[instruction.output] = v[instruction.input1] | v[instruction.input2]
		} else if instruction.gate == "XOR" {
			v[instruction.output] = v[instruction.input1] ^ v[instruction.input2]
		}
		visited[instruction.output] = struct{}{}
		if strings.HasPrefix(instruction.output, "z") {
			outputs = append(outputs, instruction.output)
		}
	}

	slices.Sort(outputs)
	var binary string
	for i := len(outputs) - 1; i >= 0; i-- {
		binary += strconv.Itoa(v[outputs[i]])
	}
	decimal, _ := strconv.ParseInt(binary, 2, 64)

	return int(decimal)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	values := map[string]int{}
	instructions := []Instruction{}
	var xBinary, yBinary string

	for _, value := range strings.Split(split[0], "\n") {
		valueSplit := strings.Split(value, ": ")
		num, _ := strconv.Atoi(valueSplit[1])
		values[valueSplit[0]] = num
		if strings.HasPrefix(valueSplit[0], "x") {
			xBinary = valueSplit[1] + xBinary
		}
		if strings.HasPrefix(valueSplit[0], "y") {
			yBinary = valueSplit[1] + yBinary
		}
	}

	for _, instruction := range strings.Split(split[1], "\n") {
		instructionSplit := strings.Split(instruction, " ")
		instructions = append(instructions, Instruction{
			input1: instructionSplit[0],
			input2: instructionSplit[2],
			gate:   instructionSplit[1],
			output: instructionSplit[4],
		})
	}

	fmt.Println("Part 1:", solve(instructions, values))

	errors := []Instruction{}
	for _, instruction := range instructions {
		checkForErrors(instruction, instructions, &errors)
	}

	output := []string{}
	for _, error := range errors {
		output = append(output, error.output)
	}

	slices.Sort(output)
	fmt.Println("Part 2:", strings.Join(output, ","))
}

func checkForErrors(instruction Instruction, instructions []Instruction, errors *[]Instruction) {
	if strings.HasPrefix(instruction.output, "z") {
		if instruction.gate != "XOR" && instruction.output != "z45" {
			*errors = append(*errors, instruction)
		}
	} else if (!slices.Contains([]byte{'x', 'y'}, instruction.input1[0]) && !slices.Contains([]byte{'x', 'y'}, instruction.input2[0])) && instruction.gate == "XOR" {
		*errors = append(*errors, instruction)
	} else if (slices.Contains([]byte{'x', 'y'}, instruction.input1[0]) && slices.Contains([]byte{'x', 'y'}, instruction.input2[0])) && instruction.gate == "XOR" {
		for _, ins := range instructions {
			if (ins.input1 == instruction.output || ins.input2 == instruction.output) && ins.gate == "XOR" {
				return
			}
		}
		*errors = append(*errors, instruction)
	} else if instruction.gate == "AND" {
		for _, ins := range instructions {
			if slices.Contains([]string{"x00", "y00"}, instruction.input1) && slices.Contains([]string{"x00", "y00"}, instruction.input2) {
				return
			}
			if (ins.input1 == instruction.output || ins.input2 == instruction.output) && ins.gate == "OR" {
				return
			}
		}
		*errors = append(*errors, instruction)
	}
}
