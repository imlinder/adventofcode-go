package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1: %v \n", part1(string(input)))
}

func part1(input string) int {
	sum := 0
	r1 := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := r1.FindAllString(input, -1)

	for _, match := range matches {
		nums := match[4 : len(match)-1]
		numsSlice := strings.Split(nums, ",")

		num1, err := strconv.Atoi(numsSlice[0])
		if err != nil {
			panic(err)
		}

		num2, err := strconv.Atoi(numsSlice[1])
		if err != nil {
			panic(err)
		}

		sum += num1 * num2
	}

	return sum
}
