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
	found := 0

	for _, l := range lines {
		if len(l) < 1 {
			continue
		}
		ranges := strings.Split(l, ",")
		if contains(ranges[0], ranges[1]) {
			found += 1
		}
	}

	return found
}

func contains(r1 string, r2 string) bool {
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
