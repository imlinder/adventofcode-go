package main

import (
	"fmt"
	"github.com/imlinder/adventofcode-go/util"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadFile("input.txt")
	fmt.Printf("Part 1: %v \n", part1(lines))
}

func part1(lines []string) int {
	sum := 0

	var left []int
	var right []int

	for _, l := range lines {
		split := strings.Fields(l)
		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[1])
		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		sum += max(diff, -diff)
	}

	return sum
}
