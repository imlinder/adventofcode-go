package main

import (
	"fmt"
	"github.com/imlinder/adventofcode-go/util"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadFile("input.txt")
	part1, part2 := run(lines)
	fmt.Printf("Part 1: %v \n", part1)
	fmt.Printf("Part 2: %v \n", part2)
}

func run(lines []string) (int, int) {
	contains := 0
	overlaps := 0

	for _, l := range lines {
		if len(l) < 1 {
			continue
		}
		ranges := strings.Split(l, ",")
		if rangesContains(ranges[0], ranges[1]) {
			contains += 1
		}
		if rangesOverlaps(ranges[0], ranges[1]) {
			overlaps += 1
		}
	}

	return contains, overlaps
}

func rangesContains(r1 string, r2 string) bool {
	s1, s2 := strings.Split(r1, "-"), strings.Split(r2, "-")

	l1, _ := strconv.Atoi(s1[0])
	h1, _ := strconv.Atoi(s1[1])
	l2, _ := strconv.Atoi(s2[0])
	h2, _ := strconv.Atoi(s2[1])

	if l1 >= l2 && h1 <= h2 {
		return true
	}

	if l2 >= l1 && h2 <= h1 {
		return true
	}

	return false
}

func rangesOverlaps(r1 string, r2 string) bool {
	s1, s2 := strings.Split(r1, "-"), strings.Split(r2, "-")

	l1, _ := strconv.Atoi(s1[0])
	h1, _ := strconv.Atoi(s1[1])
	l2, _ := strconv.Atoi(s2[0])
	h2, _ := strconv.Atoi(s2[1])

	if h1 >= l2 && h1 <= h2 {
		return true
	}

	if h2 >= l1 && h2 <= h1 {
		return true
	}

	return false
}
