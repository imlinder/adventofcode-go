package main

import (
	"fmt"
	"github.com/imlinder/adventofcode-go/util"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadFile("input.txt")
	fmt.Printf("Part 1: %v \n", part1(lines))
}

func part1(lines []string) int {

	sum := 0

	for _, l := range lines {
		if isSafe(strings.Fields(l)) {
			sum++
		} else {
		}
	}

	return sum
}

func isSafe(rule []string) bool {

	irule := SliceToInts(rule)
	prev := irule[0]

	isIncreasing := irule[0] < irule[1]

	for i := 1; i < len(irule); i++ {
		curr := irule[i]
		diff := curr - prev

		if curr == prev {
			return false
		}

		if max(diff, -diff) > 3 {
			return false
		}

		if isIncreasing && curr < prev {
			return false
		}

		if !isIncreasing && curr > prev {
			return false
		}

		prev = curr
	}

	return true
}

func SliceToInts(sa []string) []int {
	n := []int{}
	for _, s := range sa {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		n = append(n, i)
	}

	return n
}
