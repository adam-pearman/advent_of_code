package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mixAndPrune(secret, value int) int {
	return (secret ^ value) % 16777216
}

func getSecret(secret, iterations int) int {
	for i := 0; i < iterations; i++ {
		secret = mixAndPrune(secret, secret*64)
		secret = mixAndPrune(secret, secret/32)
		secret = mixAndPrune(secret, secret*2048)
	}
	return secret
}

func getSequence(secret, iterations int) map[[4]int]int {
	sequence := make(map[[4]int]int)
	prev := secret
	digits := make([]int, 0)
	diffs := make([]int, 0)
	for i := 0; i < iterations; i++ {
		actual := getSecret(prev, 1)
		digits = append(digits, actual%10)
		diffs = append(diffs, actual%10-prev%10)
		prev = actual
	}
	for i := 3; i < len(diffs); i++ {
		seq := [4]int{diffs[i-3], diffs[i-2], diffs[i-1], diffs[i]}
		if _, ok := sequence[seq]; !ok {
			sequence[seq] = digits[i]
		}
	}
	return sequence
}

func main() {
	input, _ := os.ReadFile("input.txt")
	secrets := strings.Split(strings.TrimSpace(string(input)), "\n")

	secretSum := 0
	bananas := make(map[[4]int]int)
	for _, secret := range secrets {
		secret, _ := strconv.Atoi(secret)
		secretSum += getSecret(secret, 2000)
		sequence := getSequence(secret, 2000)
		for key, value := range sequence {
			bananas[key] += value
		}
	}
	fmt.Println("Part 1:", secretSum)

	maxValue := 0
	for _, count := range bananas {
		if count > maxValue {
			maxValue = count
		}
	}
	fmt.Println("Part 2:", maxValue)
}
